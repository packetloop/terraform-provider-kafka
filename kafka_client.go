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
}

func(info *KafkaTopicInfo) exists() bool {
  return info != nil && info.PartitionsCount > 0 && info.ReplicationFactor > 0
}

func (client *KafkaManagingClient) createTopic(name string, partitions int, replicas int) (*KafkaTopicInfo, error) {
  cmd := exec.Command(
    client.TopicScript,
    "--zookeeper", client.Zookeeper,
    "--create", "--topic", name,
    "--partitions", strconv.Itoa(partitions),
    "--replication-factor", strconv.Itoa(replicas))

  out, err := cmd.Output()

  if err != nil {
    kafkaError := readError(string(out))
    if (kafkaError != nil) { return nil, kafkaError }
    return nil, err
  }

  strOut := strings.TrimSpace(string(out))
  logrus.Info(strOut)
  log.Println(strOut)

  if strOut == fmt.Sprintf("Created topic \"%s\".", name) {
    res := &KafkaTopicInfo{
      PartitionsCount: partitions,
      ReplicationFactor: replicas,
    }
    return res, nil
  }

  return nil, fmt.Errorf("Unable to parse results from kafka, there is maybe something wrong: %s", strOut)
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
  partsR, _ := regexp.Compile("PartitionCount:(\\d+).+ReplicationFactor:(\\d+)")
  pRes := partsR.FindStringSubmatch(txt)
  if len(pRes) != 3 {
    return nil, fmt.Errorf("Unable to determine topic's partitions count")
  }

  pCount, pcErr := strconv.Atoi(pRes[1])
  if pcErr != nil {
    return nil, fmt.Errorf("Unable to read topic's partition count: " + pcErr.Error())
  }

  rCount, rErr := strconv.Atoi(pRes[2])
  if rErr != nil {
    return nil, fmt.Errorf("Unable to read topic's partition count: " + rErr.Error())
  }

  info := &KafkaTopicInfo{
    PartitionsCount: pCount,
    ReplicationFactor: rCount,
  }

  return info, nil
}


func createRequest(id string, config interface{}) error {
  // TODO POST to api with payload
  return nil
}

func createDeploy(id string, config interface{}) error {
  // TODO POST to api with payload
  return nil
}

func getRequest(id string, meta interface{}) error {
  // TODO GET to api, parse response into request/deploy objects and return
  return nil
}

func waitForRequest(id string, status string, meta interface{}) error {

  activeState := false

  // TODO actually parse response and compare status

  // var endpoint = meta.(*Conf).endpoint + "/requests/request/" + id

  for activeState == false {
    // res, err := http.Get(endpoint)

    // if err != nil {
    //  return err
    // }

    // defer r.Body.Close()
    // decoder := json.NewDecoder(res.Body)

    // activeState = *res.Table.TableStatus == "ACTIVE"

    // // Wait for a few seconds
    // if !activeState {
    //  log.Printf("[DEBUG] Sleeping for 5 seconds for table to become active")
    //  time.Sleep(5 * time.Second)
    // }
  }

  return nil
}

func deleteRequest(id string, meta interface{}) error {
  return nil
}


