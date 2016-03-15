package main

import (
  "log"
  "os/exec"
  "regexp"
  "strconv"
  "strings"
  "fmt"
  "github.com/hashicorp/consul/vendor/github.com/fsouza/go-dockerclient/external/github.com/Sirupsen/logrus"
)

// Client does client stuff.
type KafkaManagingClient struct {
  Zookeeper   string
  TopicScript string
}

type KafkaTopicInfo struct {
  PartitionsCount int
  ReplicationFactor int
  CleanupPolicy string
  RetentionBytes int64
  RetentionMs int64
}

func appendConf(slice []string, name string, value string) []string {
  return append(slice, "--config", name + "=" + value)
}

func(conf *KafkaTopicInfo) kafkaTopicConfigOpts() []string {
  var parms = []string{}
  if (conf.CleanupPolicy != "") { parms = appendConf(parms, "cleanup.policy", conf.CleanupPolicy ) }
  if (conf.RetentionBytes > -1) { parms = appendConf(parms, "retention.bytes", strconv.FormatInt(conf.RetentionBytes, 10))}
  if (conf.RetentionMs > -1)    { parms = appendConf(parms, "retention.ms", strconv.FormatInt(conf.RetentionMs, 10))}

  return parms
}

func(info *KafkaTopicInfo) exists() bool {
  return info != nil && info.PartitionsCount > 0 && info.ReplicationFactor > 0
}

func newKafkaTopicInfo(partitions int, replicas int) *KafkaTopicInfo {
  return &KafkaTopicInfo{
    PartitionsCount:   partitions,
    ReplicationFactor: replicas,
    CleanupPolicy:     "",
    RetentionBytes:    -1,
    RetentionMs:       -1,
  }
}

func (client *KafkaManagingClient) deleteTopic(name string) error {
  cmd := exec.Command(
    client.TopicScript,
    "--zookeeper", client.Zookeeper,
    "--delete", "--topic", name)

  out, err := cmd.Output()
  if err != nil {
    kafkaError := readError(string(out))
    if (kafkaError != nil) { return kafkaError }
    return err
  }

  strOut := strings.TrimSpace(string(out))
  if strings.Contains(strOut, "marked for deletion") {
    return nil
  }

  return fmt.Errorf("Was not able to confirm that topic %s was marked for deletion. Something is wrong", name)
}

func (client *KafkaManagingClient) createTopic(name string, conf *KafkaTopicInfo) error {
  var params = []string {
    "--zookeeper", client.Zookeeper,
    "--create", "--topic", name,
    "--partitions", strconv.Itoa(conf.PartitionsCount),
    "--replication-factor", strconv.Itoa(conf.ReplicationFactor),
  }

  confOpts := conf.kafkaTopicConfigOpts()
  params = append(params, confOpts...)

  log.Println("[DEBUG] Will execute %v", params)

  cmd := exec.Command(
    client.TopicScript,
    params...
    )

  out, err := cmd.Output()

  if err != nil {
    kafkaError := readError(string(out))
    if (kafkaError != nil) { return kafkaError }
    return err
  }

  strOut := strings.TrimSpace(string(out))
  logrus.Info(strOut)
  log.Println(strOut)

  if strOut == fmt.Sprintf("Created topic \"%s\".", name) {
    return nil
  }

  return fmt.Errorf("Unable to parse results from kafka, there is maybe something wrong: %s", strOut)
}

func (client *KafkaManagingClient) describeTopic(name string) (*KafkaTopicInfo, error) {
  cmd := exec.Command(client.TopicScript, "--zookeeper", client.Zookeeper, "--describe", "--topic", name)

  out, err := cmd.Output()
  if err != nil { return nil, err }

  strOut := strings.TrimSpace(string(out[:len(out)]))

  //does not exist
  if strOut == "" { return nil, nil }

  return readTopicInfo(strOut)
}

func readError(txt string) error {
  errorR, _ := regexp.Compile("^Error .+")
  err := strings.TrimSpace(errorR.FindString(txt))

  if err == "" { return nil }
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
    if !strings.Contains(e, "=") { continue }
    ps := strings.SplitN(e, "=", 2)
    confOpts[ps[0]] = ps[1]
  }

  info := &KafkaTopicInfo{
    PartitionsCount:   pCount,
    ReplicationFactor: rCount,
    CleanupPolicy:     getOrDefaultStr(confOpts, "cleanup.policy", ""),
    RetentionBytes:    getOrDefaultInt(confOpts, "retention.bytes", -1),
    RetentionMs:       getOrDefaultInt(confOpts, "retention.ms", -1),
  }

  return info, nil
}

func getOrDefaultStr(m map[string]string, key string, def string) string {
  if v, ok := m[key]; ok { return v }
  return def
}

func getOrDefaultInt(m map[string]string, key string, def int64) int64 {
  if v, ok := m[key]; ok {
    if i, pok := strconv.ParseInt(v, 10, 0); pok == nil { return i }
  }
  return def
}

