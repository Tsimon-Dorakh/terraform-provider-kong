# kong_consumer Data Source

Provides details about a specific Kong consumer.
For more information on the parameters [see the Kong Consumer](https://getkong.org/docs/1.0.x/admin-api/#consumer-object).

## Example Usage

```hcl
data "kong_consumer" "consumer" {
    username  = "User1"
}
```

## Argument Reference

* `username` - (Required) The username to use

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - An id of the consumer
* `custom_id` - A custom id of the consumer