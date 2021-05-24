package kong

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKongConsumer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceKongConsumerRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceKongConsumerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	consumerClient := meta.(*config).adminClient.Consumers()
	consumer, err := consumerClient.GetByUsername(d.Get("username").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	if consumer == nil {
		d.SetId("")
	} else {
		d.SetId(consumer.Id)
		d.Set("id", consumer.Id)
		d.Set("username", consumer.Username)
		d.Set("custom_id", consumer.CustomId)
	}

	return diags
}
