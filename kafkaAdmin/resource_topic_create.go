package kafkaadmin

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	kafka "github.com/packetloop/go-kafkaesque"
)

func resourceKafkaTopic() *schema.Resource {
	return &schema.Resource{
		Create: resourceKafkaTopicCreate,
		Read:   resourceKafkaTopicRead,
		Exists: resourceKafkaTopicExists,
		Update: resourceKafkaTopicUpdate,
		Delete: resourceKafkaTopicDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Kafka topic name",
			},
			"partitions": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "number of partitions for the topic",
				Default:     1,
			},
			"replication_factor": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "the replication factor for the topic",
				Default:     1,
			},
			"retention_ms": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "the retention period in milliseconds for the topic",
				Default:     -1,
			},
			"cleanup_policy": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "the clean up policy for the topic, for example compaction",
				Default:      "compact",
				ValidateFunc: validateCleanupPolicy,
			},
			"segment_bytes": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "the segment file size for the log",
				Default:     1073741824,
			},
			"min_insync_replicas": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "the minimum number of insync replicas",
				Default:     1,
			},
			"segment_ms": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "the time after which Kafka will force the log to roll",
				Default:     604800000,
			},
			"retention_bytes": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "the retention bytes for the topic",
				Default:     -1,
			},
		},
	}
}

func resourceKafkaTopicCreate(d *schema.ResourceData, m interface{}) error {

	id := d.Get("name").(string)
	d.SetId(id)
	return createRequest(d, m)
}

func resourceKafkaTopicExists(d *schema.ResourceData, m interface{}) (b bool, e error) {
	// Exists - This is called to verify a resource still exists. It is called prior to Read,
	// and lowers the burden of Read to be able to assume the resource exists.
	client := clientConn(m)
	_, err := client.GetTopic(d.Id())
	if err != nil {
		return false, fmt.Errorf("TOPIC '%v' DOES NOT EXIST: %v", d.Id(), err)
	}
	return true, nil

}

func createRequest(d *schema.ResourceData, m interface{}) error {
	id := strings.ToLower(d.Get("name").(string))
	partitions := d.Get("partitions").(string)
	replicationFactor := d.Get("replication_factor").(string)
	retentionMs := d.Get("retention_ms").(string)
	cleanupPolicy := d.Get("cleanup_policy").(string)
	segmentBytes := d.Get("segment_bytes").(string)
	retentionBytes := d.Get("retention_bytes").(string)
	segmentMs := d.Get("segment_ms").(string)
	minInsyncReplicas := d.Get("min_insync_replicas").(string)

	log.Printf("[TRACE] creating kafka topic '%s'...", id)
	client := clientConn(m)
	t := kafka.NewTopic(id).
		SetReplicationFactor(replicationFactor).
		SetPartitions(partitions).
		BuildTopic()
	t.Config = &kafka.Config{
		RetentionMs:       retentionMs,
		SegmentBytes:      segmentBytes,
		CleanupPolicy:     cleanupPolicy,
		MinInsyncReplicas: minInsyncReplicas,
		RetentionBytes:    retentionBytes,
		SegmentMs:         segmentMs,
	}
	resp, err := client.CreateTopic(t)
	if err != nil {
		log.Printf("[DEBUG] Error Response %v", err)
	}

	return checkResponse(d, m, resp, err)
}

func checkResponse(d *schema.ResourceData, m interface{}, r kafka.Response, err error) error {
	log.Printf("[TRACE] Create Topic %v", r)
	if err != nil {
		return fmt.Errorf("CREATE TOPIC '%s' ERROR: %v", d.Id(), err)
	}
	return resourceKafkaTopicRead(d, m)
}

// resourceKafkaTopicRead is called to resync the local state with the remote state.
// Terraform guarantees that an existing ID will be set. This ID should be used
// to look up the resource. Any remote data should be updated into the local data.
// No changes to the remote resource are to be made.
func resourceKafkaTopicRead(d *schema.ResourceData, m interface{}) error {
	client := clientConn(m)

	r, err := client.GetTopic(d.Id())
	if err != nil {
		return fmt.Errorf("GETTING TOPIC '%s' ERROR: %v", d.Id(), err)
	}

	// Unfortunately get topics does not return name of topic, only its config params.
	d.Set("name", d.Id())
	d.Set("partitions", r.GetPartitions())
	d.Set("replication_factor", r.GetReplicationFactor())
	d.Set("retention_ms", r.GetRetentionMs())
	d.Set("cleanup_policy", r.GetCleanupPolicy())
	d.Set("segment_bytes", r.GetSegmentBytes())
	d.Set("min_insync_replicas", r.GetMinInSyncReplicas())
	d.Set("retention_bytes", r.GetRetentionBytes())
	d.Set("segment_ms", r.GetSegmentMs())
	return nil

}

func resourceKafkaTopicUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	if d.HasChange("partitions") ||
		d.HasChange("retention_ms") ||
		d.HasChange("cleanup_policy") ||
		d.HasChange("segment_bytes") ||
		d.HasChange("segment_ms") ||
		d.HasChange("min_insync_replicas") ||
		d.HasChange("retention_bytes") {
		log.Printf("[TRACE] update existing topic (%s) success", d.Id())
		d.Partial(false)

		return updateRequest(d, m)
	}
	return nil
}

func updateRequest(d *schema.ResourceData, m interface{}) error {
	id := strings.ToLower(d.Get("name").(string))
	partitions := d.Get("partitions").(string)
	replicationFactor := d.Get("replication_factor").(string)
	retentionMs := d.Get("retention_ms").(string)
	cleanupPolicy := d.Get("cleanup_policy").(string)
	segmentBytes := d.Get("segment_bytes").(string)
	retentionBytes := d.Get("retention_bytes").(string)
	segmentMs := d.Get("segment_ms").(string)
	minInsyncReplicas := d.Get("min_insync_replicas").(string)

	log.Printf("[TRACE] updating kafka topic '%s'...", id)

	client := clientConn(m)

	t := kafka.NewTopic(id).
		SetReplicationFactor(replicationFactor).
		SetPartitions(partitions).
		BuildTopic()
	t.Config = &kafka.Config{
		RetentionMs:       retentionMs,
		SegmentBytes:      segmentBytes,
		CleanupPolicy:     cleanupPolicy,
		MinInsyncReplicas: minInsyncReplicas,
		RetentionBytes:    retentionBytes,
		SegmentMs:         segmentMs,
	}

	resp, err := client.UpdateTopic(t)
	if err != nil {
		log.Printf("[DEBUG] Error Response %v", err)
	}

	return checkResponse(d, m, resp, err)
}

func resourceKafkaTopicDelete(d *schema.ResourceData, m interface{}) error {
	a := deleteRequest(d.Id())
	return a(d, m)
}

func deleteRequest(id string) (f func(d *schema.ResourceData, m interface{}) error) {
	return func(d *schema.ResourceData, m interface{}) error {
		client := clientConn(m)
		// Return 'Ok' when successful. Otherwise, this throws an error. Hence,
		// we can safely ignore this.
		_, err := client.DeleteTopic(id)
		if err != nil {
			return err
		}
		d.SetId("")
		return nil
	}
}
