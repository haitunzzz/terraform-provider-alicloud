---
subcategory: "KMS"
layout: "alicloud"
page_title: "Alicloud: alicloud_kms_instance"
description: |-
  Provides a Alicloud KMS Instance resource.
---

# alicloud_kms_instance

Provides a KMS Instance resource. 

For information about KMS Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/zh/key-management-service/latest/kms-instance-management).

-> **NOTE:** Available since v1.210.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

data "alicloud_vpcs" "default" {
  name_regex = "^default-NODELETING$"
  cidr_block = "172.16.0.0/16"
}
data "alicloud_vswitches" "default" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = "cn-hangzhou-h"
}

resource "alicloud_kms_instance" "default" {
  product_version = "3"
  vpc_id          = data.alicloud_vpcs.default.ids.0
  zone_ids = [
    "cn-hangzhou-h",
    "cn-hangzhou-g"
  ]
  vswitch_ids = [
    data.alicloud_vswitches.default.ids.0
  ]
  vpc_num    = "1"
  key_num    = "1000"
  secret_num = "0"
  spec       = "1000"
}

# Save Instance's CA certificate chain to a local file
# resource "local_file" "ca_certificate_chain_pem" {
#   content  = alicloud_kms_instance.default.ca_certificate_chain_pem
#   filename = "ca.pem"
# }
```

## Argument Reference

The following arguments are supported:
* `bind_vpcs` - (Optional) Aucillary VPCs used to access this KMS instance. See [`bind_vpcs`](#bind_vpcs) below.
* `force_delete_without_backup` - (Optional, Available since v1.223.2) Whether to force deletion even without backup.
* `key_num` - (Optional) Maximum number of stored keys. The attribute is valid when the attribute `payment_type` is `Subscription`.
* `log` - (Optional, Computed) Instance Audit Log Switch. The attribute is valid when the attribute `payment_type` is `Subscription`.
* `log_storage` - (Optional, Computed) Instance log capacity. The attribute is valid when the attribute `payment_type` is `Subscription`.
* `payment_type` - (Optional, ForceNew) Payment type, valid values:  `Subscription`: Prepaid. `PayAsYouGo`: Postpaid, available since v1.223.2.
* `period` - (Optional) Purchase cycle, in months. The attribute is valid when the attribute `payment_type` is `Subscription`.
* `product_version` - (Optional) KMS Instance commodity type (software/hardware).
* `renew_period` - (Optional) Automatic renewal period, in months. The attribute is valid when the attribute `payment_type` is `Subscription`.
* `renew_status` - (Optional) Renewal options. Valid values: `AutoRenewal`, `ManualRenewal`. The attribute is valid when the attribute `payment_type` is `Subscription`.
* `secret_num` - (Optional) Maximum number of Secrets. The attribute is valid when the attribute `payment_type` is `Subscription`.
* `spec` - (Optional) The computation performance level of the KMS instance. The attribute is valid when the attribute `payment_type` is `Subscription`.
* `vpc_id` - (Required, ForceNew) Instance VPC id.
* `vpc_num` - (Optional) The number of managed accesses. The maximum number of VPCs that can access this KMS instance. The attribute is valid when the attribute `payment_type` is `Subscription`.
* `vswitch_ids` - (Required, ForceNew) Instance bind vswitches.
* `zone_ids` - (Required, ForceNew) zone id.

### `bind_vpcs`

The bind_vpcs supports the following:
* `region_id` - (Optional) region id.
* `vswitch_id` - (Optional) vswitch id.
* `vpc_id` - (Optional) VPC ID.
* `vpc_owner_id` - (Optional) VPC owner root user ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `ca_certificate_chain_pem` - KMS instance certificate chain in PEM format.
* `create_time` - The creation time of the resource.
* `instance_name` - The name of the resource.
* `status` - Instance status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 60 mins) Used when create the Instance.
* `delete` - (Defaults to 10 mins) Used when delete the Instance.
* `update` - (Defaults to 60 mins) Used when update the Instance.

## Import

KMS Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_kms_instance.example <id>
```