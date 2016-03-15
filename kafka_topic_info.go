package main
import "strconv"

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

func addConf(slice []string, name string, value string) []string {
  return append(slice, "--add-config", name + "=" + value)
}

func removeConf(slice []string, name string) []string {
  return append(slice, "--delete-config", name)
}

func setConf(slice []string, name string, value string, nullValue string) []string {
  if nullValue == value {
    return removeConf(slice, name)
  } else {
    return addConf(slice, name, value)
  }
}

func setConfInt(slice []string, name string, value int, nullValue int) []string {
  return setConf(slice, name, strconv.Itoa(value), strconv.Itoa(nullValue))
}

func setConfInt64(slice []string, name string, value int64, nullValue int64) []string {
  return setConf(slice, name, strconv.FormatInt(value, 10), strconv.FormatInt(nullValue, 10))
}

func (conf *KafkaTopicInfo) alterTopicConfigOpts() []string {
  var parms = []string{}
  parms = setConf      (parms, "cleanup.policy",  conf.CleanupPolicy,  "")
  parms = setConfInt64 (parms, "retention.bytes", conf.RetentionBytes, -1)
  parms = setConfInt64 (parms, "retention.ms",    conf.RetentionMs,    -1)

  return parms
}

func(conf *KafkaTopicInfo) createTopicConfigOpts() []string {
  var parms = []string{}
  if (conf.CleanupPolicy != "") { parms = appendConf(parms, "cleanup.policy", conf.CleanupPolicy ) }
  if (conf.RetentionBytes > -1) { parms = appendConf(parms, "retention.bytes", strconv.FormatInt(conf.RetentionBytes, 10))}
  if (conf.RetentionMs > -1)    { parms = appendConf(parms, "retention.ms", strconv.FormatInt(conf.RetentionMs, 10))}

  return parms
}

func(info *KafkaTopicInfo) exists() bool {
  return info != nil && info.PartitionsCount > 0 && info.ReplicationFactor > 0
}