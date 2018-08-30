package kafkaadmin

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	kafka "github.com/packetloop/go-kafkaesque"
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
						"kafka_topic.foo", "name", "mytopic"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "partitions", "2"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "replication_factor", "3"),
					/* Not yet implemented
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "cleanup_policy", "1073741824"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "retention_ms", "300000"),
					resource.TestCheckResourceAttr(
						"kafka_topic.foo", "segment_bytes", "1073741824"),
					*/
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

/* Not yet implemented
retention_ms = 300000
cleanup_policy = "compact"
segment_bytes = 1073741824
*/
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

		// If topic exist, returns error nil.
		if _, err := client.GetTopic(id); err != nil {
			return fmt.Errorf("ERROR TOPIC '%s' DOES NOT EXIST: %v", id, err)
		}
	}
	return nil
}
