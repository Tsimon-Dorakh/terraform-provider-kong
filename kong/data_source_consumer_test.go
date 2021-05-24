package kong

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceKongConsumer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testConsumerDataSourceCreateConsumerConfig,
			},
			{
				Config: testConsumerDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.kong_consumer.consumer_data_source", "username", "User1"),
					resource.TestCheckResourceAttr("data.kong_consumer.consumer_data_source", "custom_id", "123"),
				),
			},
		},
	})
}

const testConsumerDataSourceCreateConsumerConfig = `
resource "kong_consumer" "consumer" {
	username  = "User1"
	custom_id = "123"
}
`

const testConsumerDataSourceConfig = `
data "kong_consumer" "consumer_data_source" {
	username  = "User1"
}
`
