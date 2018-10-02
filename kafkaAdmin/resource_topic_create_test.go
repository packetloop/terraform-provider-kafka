package kafkaadmin

import (
	"fmt"
	"testing"

	kafka "github.com/comozo/go-kafkaesque"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccKafkaAdminTopicCreate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		PreCheck:                  func() { testAccPreCheck(t) },
		Providers:                 testAccProviders,
		CheckDestroy:              testCheckTopicDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKafkaTopicCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTopicExists("kafka_topic.foo"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "name", "foo"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "partitions", "2"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "replication_factor", "3"),
				),
			},
		},
	})
}

const testAccCheckKafkaTopicCreate = `
resource "kafka_topic" "foo" {
  name = "mytopic"
  partitions = 2
  replication_factor = 3
}
`

func TestAccKafkaAdminTopicCreateWithConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		PreCheck:                  func() { testAccPreCheck(t) },
		Providers:                 testAccProviders,
		CheckDestroy:              testCheckTopicDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKafkaTopicCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTopicExists("kafka_topic.foobar"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foobar", "name", "mytopicconfig"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foobar", "partitions", "2"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foobar", "replication_factor", "3"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foobar", "cleanup_policy", "delete"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foobar", "retention_ms", "300000"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foobar", "segment_bytes", "10737418"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foobar", "segment_ms", "600000"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foobar", "retention_bytes", "100000"),
				),
			},
		},
	})
}

const testAccCheckKafkaTopicCreateConfig = `
resource "kafka_topic" "foobar" {
  name = "mytopicconfig"
  partitions = 2
  replication_factor = 3
  retention_ms = 300000
  cleanup_policy = "delete"
  segment_bytes = 10737418
  min_insync_replicas = 2
  retention_bytes = 100000
  segment_ms = 600000
}
`

func TestAccKafkaAdminTopicUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		PreCheck:                  func() { testAccPreCheck(t) },
		Providers:                 testAccProviders,
		CheckDestroy:              testCheckTopicDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckKafkaTopicCreate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTopicExists("kafka_topic.foo"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "name", "mytopic"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "partitions", "2"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "replication_factor", "3"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "retention_ms", "-1"),
				),
			},
			{
				Config: testAccCheckKafkaTopicUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTopicExists("kafka_topic.foo"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "name", "mytopic"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "partitions", "2"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "replication_factor", "3"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "retention_ms", "100000"),
				),
			},
		},
	})
}

const testAccCheckKafkaTopicUpdate = `
resource "kafka_topic" "foo" {
  name = "mytopic"
  partitions = 2
  replication_factor = 3
  retention_ms = 100000
}
`

func testAccCheckTopicExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Conn).sclient
		return TopicExistsHelper(s, client)
	}
}

func testCheckTopicDestroy(state *terraform.State) error {
	for _, res := range state.RootModule().Resources {
		if res.Type != "kafka_topic" {
			continue
		}

		topicName := res.Primary.ID

		client := testAccProvider.Meta().(*Conn).sclient
		_, err := client.GetTopic(topicName)
		// If err is not nil, it means topic still exist.
		if err == nil {
			return fmt.Errorf("ERROR DESTROY TOPIC '%s': %v", topicName, err)
		}
		return nil
	}
	return nil
}

func TopicExistsHelper(s *terraform.State, client *kafka.Client) error {
	for _, r := range s.RootModule().Resources {
		id := r.Primary.ID

		_, err := client.GetTopic(id)
		status, _ := errorHelper(err)
		if status.state == Exists {
			return fmt.Errorf("ERROR TOPIC '%s': %v", id, err)
		}
	}
	return nil
}
