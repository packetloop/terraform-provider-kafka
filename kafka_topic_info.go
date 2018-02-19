package main

import "strconv"

type KafkaTopicInfo struct {
	PartitionsCount   int
	ReplicationFactor int
	CleanupPolicy     string
	RetentionBytes    int64
	RetentionMs       int64
	SegmentBytes      int64
	SegmentMs         int64
	PartitionsCountChanged   bool
	ReplicationFactorChanged bool
	CleanupPolicyChanged     bool
	RetentionBytesChanged    bool
	RetentionMsChanged       bool
	SegmentBytesChanged      bool
	SegmentMsChanged         bool
}

type ConfMods struct {
	ConfDeletions map[string]string
	ConfAdditions map[string]string
}

func appendConf(slice []string, name string, value string) []string {
	return append(slice, "--config", name+"="+value)
}

func setConf(confMods ConfMods, name string, changed bool, value string, nullValue string) {
	if changed {
		if nullValue == value {
			confMods.ConfDeletions[name] = value
		} else {
			confMods.ConfAdditions[name] = value
		}
	}
}

func setConfInt(confMods ConfMods, name string, changed bool, value int, nullValue int) {
	setConf(confMods, name, changed, strconv.Itoa(value), strconv.Itoa(nullValue))
}

func setConfInt64(confMods ConfMods, name string, changed bool, value int64, nullValue int64) {
	setConf(confMods, name, changed, strconv.FormatInt(value, 10), strconv.FormatInt(nullValue, 10))
}

func makeConfMods() ConfMods {
	confMods := ConfMods {
		ConfDeletions: make(map[string]string),
		ConfAdditions: make(map[string]string)}

	return confMods
}

func writeConfMods(slice []string, conf *ConfMods) []string {
	if (len(conf.ConfAdditions) > 0) {
		arg := ""
		delim := ""

		for k, v := range conf.ConfAdditions {
			arg += delim + k + "=" + v
			delim = ","
		}

		slice = append(slice, "--add-config")
		slice = append(slice, arg)
	}

	for k, _ := range conf.ConfDeletions {
		slice = append(slice, "--delete-config", k)
	}

	return slice
}

func (conf *KafkaTopicInfo) alterTopicConfigOpts() []string {
	var parms = []string{}

	confMods := ConfMods {
		ConfDeletions: make(map[string]string),
		ConfAdditions: make(map[string]string)}

	setConf     (confMods, "cleanup.policy" , conf.CleanupPolicyChanged, 	conf.CleanupPolicy  , "")
	setConfInt64(confMods, "retention.bytes", conf.RetentionBytesChanged,	conf.RetentionBytes , -1)
	setConfInt64(confMods, "retention.ms"   , conf.RetentionMsChanged,   	conf.RetentionMs    , -1)
	setConfInt64(confMods, "retention.ms"   , conf.RetentionMsChanged,   	conf.RetentionMs    , -1)
	setConfInt64(confMods, "segment.bytes"  , conf.SegmentBytesChanged,  	conf.SegmentBytes   , -1)
	setConfInt64(confMods, "segment.ms"     , conf.SegmentMsChanged,     	conf.SegmentMs      , -1)

	parms = writeConfMods(parms, &confMods)

	// parms = append(parms, "--moo")

	return parms
}

func (conf *KafkaTopicInfo) createTopicConfigOpts() []string {
	var parms = []string{}
	if conf.CleanupPolicy != "" {
		parms = appendConf(parms, "cleanup.policy", conf.CleanupPolicy)
	}
	if conf.RetentionBytes > -1 {
		parms = appendConf(parms, "retention.bytes", strconv.FormatInt(conf.RetentionBytes, 10))
	}
	if conf.RetentionMs > -1 {
		parms = appendConf(parms, "retention.ms", strconv.FormatInt(conf.RetentionMs, 10))
	}

	if conf.SegmentBytes > -1 {
		parms = appendConf(parms, "segment.bytes", strconv.FormatInt(conf.SegmentBytes, 10))
	}

	if conf.SegmentMs > -1 {
		parms = appendConf(parms, "segment.ms", strconv.FormatInt(conf.SegmentMs, 10))
	}

	return parms
}

func (info *KafkaTopicInfo) exists() bool {
	return info != nil && info.PartitionsCount > 0 && info.ReplicationFactor > 0
}
