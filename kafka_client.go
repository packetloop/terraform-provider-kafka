package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// KafkaManagingClient does client stuff.
type KafkaManagingClient struct {
	Zookeeper    string
	TopicScript  string
	ConfigScript string
}

func (client *KafkaManagingClient) alterTopicPartitions(name string, partitions int) error {
	log.Printf("Update partitions count for topic '%s' to %d", name, partitions)
	cmd := exec.Command(
		client.TopicScript,
		"--zookeeper", client.Zookeeper,
		"--alter", "--topic", name,
		"--partitions", strconv.Itoa(partitions))

	return execKafkaCommand(cmd, "Adding partitions succeeded")
}

func (client *KafkaManagingClient) alterTopicConfig(name string, conf *KafkaTopicInfo) error {
	var params = []string{
		"--zookeeper", client.Zookeeper,
		"--entity-type", "topics",
		"--entity-name", name,
		"--alter",
	}

	confOpts := conf.alterTopicConfigOpts()
	params = append(params, confOpts...)

	log.Printf("Will update configs for topic %s: %v", name, confOpts)
	cmd := exec.Command(client.ConfigScript, params...)

	return execKafkaCommand(cmd, "Completed Updating config for entity: topic")
}

func (client *KafkaManagingClient) deleteTopic(name string) error {
	cmd := exec.Command(
		client.TopicScript,
		"--zookeeper", client.Zookeeper,
		"--delete", "--topic", name)

	return execKafkaCommand(cmd, "marked for deletion")
}

func (client *KafkaManagingClient) createTopic(name string, conf *KafkaTopicInfo) error {
	var params = []string{
		"--zookeeper", client.Zookeeper,
		"--create", "--topic", name,
		"--partitions", strconv.Itoa(conf.PartitionsCount),
		"--replication-factor", strconv.Itoa(conf.ReplicationFactor),
	}

	confOpts := conf.createTopicConfigOpts()
	params = append(params, confOpts...)

	log.Printf("[DEBUG] Will execute %v", params)

	cmd := exec.Command(client.TopicScript, params...)

	out, err := cmd.Output()

	if err != nil {
		kafkaError := readError(string(out))
		if kafkaError != nil {
			return kafkaError
		}
		return err
	}

	strOut := strings.TrimSpace(string(out))
	if strOut == fmt.Sprintf("Created topic \"%s\".", name) {
		return nil
	}

	return fmt.Errorf("Unable to parse results from kafka, there is maybe something wrong: %s", strOut)
}

func (client *KafkaManagingClient) describeTopic(name string) (*KafkaTopicInfo, error) {
	cmd := exec.Command(client.TopicScript, "--zookeeper", client.Zookeeper, "--describe", "--topic", name)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	strOut := strings.TrimSpace(string(out[:len(out)]))

	//does not exist
	if strOut == "" {
		fmt.Printf("Topic not found %s", name)
		return nil, nil
	}

	return readTopicInfo(strOut)
}

func readError(txt string) error {
	errorR, _ := regexp.Compile("(?m:^Error .+)")
	err := strings.TrimSpace(errorR.FindString(txt))

	if err == "" {
		return nil
	}
	return fmt.Errorf("%s", err)
}

func readTopicInfo(txt string) (*KafkaTopicInfo, error) {
	partsR, _ := regexp.Compile("PartitionCount:(\\d+).+ReplicationFactor:(\\d+).+Configs:\\s*([^\\s]+)?")
	pRes := partsR.FindStringSubmatch(txt)
	if len(pRes) != 4 {
		return nil, fmt.Errorf("Unable to determine topic's partitions count (Unexpected format)")
	}

	pCount, pcErr := strconv.Atoi(pRes[1])
	if pcErr != nil {
		return nil, fmt.Errorf("Unable to read topic's partition count: " + pcErr.Error())
	}

	rCount, rErr := strconv.Atoi(pRes[2])
	if rErr != nil {
		return nil, fmt.Errorf("Unable to read topic's partition count: " + rErr.Error())
	}

	// Read config options
	confOpts := make(map[string]string)

	for _, e := range strings.Split(pRes[3], ",") {
		if !strings.Contains(e, "=") {
			continue
		}
		ps := strings.SplitN(e, "=", 2)
		confOpts[ps[0]] = ps[1]
	}

	info := &KafkaTopicInfo{
		PartitionsCount:   pCount,
		ReplicationFactor: rCount,
		CleanupPolicy:     getOrDefaultStr(confOpts, "cleanup.policy", ""),
		RetentionBytes:    getOrDefaultInt(confOpts, "retention.bytes", -1),
		RetentionMs:       getOrDefaultInt(confOpts, "retention.ms", -1),
		SegmentMs:         getOrDefaultInt(confOpts, "segment.ms", -1),
		SegmentBytes:      getOrDefaultInt(confOpts, "segment.bytes", -1),
		MinInsyncReplicas: getOrDefaultInt(confOpts, "min.insync.replicas", -1),
	}

	return info, nil
}

func execKafkaCommand(cmd *exec.Cmd, successIfPresent string) error {
	out, err := cmd.Output()
	if err != nil {
		kafkaError := readError(string(out))
		if kafkaError != nil {
			return kafkaError
		}
		return err
	}

	strOut := strings.TrimSpace(string(out))
	if strings.Contains(strOut, successIfPresent) {
		return nil
	}

	return fmt.Errorf("Unable to execute command '%v': %s", cmd.Args, strOut)
}

func getOrDefaultStr(m map[string]string, key string, def string) string {
	if v, ok := m[key]; ok {
		return v
	}
	return def
}

func getOrDefaultInt(m map[string]string, key string, def int64) int64 {
	if v, ok := m[key]; ok {
		if i, pok := strconv.ParseInt(v, 10, 0); pok == nil {
			return i
		}
	}
	return def
}
