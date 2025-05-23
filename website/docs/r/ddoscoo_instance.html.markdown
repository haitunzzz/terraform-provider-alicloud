---
subcategory: "Anti-DDoS Pro (DdosCoo)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ddoscoo_instance"
sidebar_current: "docs-alicloud-resource-ddoscoo-instance"
description: |-
  Provides a Alicloud BGP-line Anti-DDoS Pro(DdosCoo) Instance Resource.
---

# alicloud_ddoscoo_instance

Provides a BGP-line Anti-DDoS Pro(DdosCoo) Instance resource.

For information about BGP-line Anti-DDoS Pro(DdosCoo) Instance and how to use it, see [What is Anti-DDoS Pro Instance](https://www.alibabacloud.com/help/en/ddos-protection/latest/create-an-anti-ddos-pro-or-anti-ddos-premium-instance-by-calling-an-api-operation).

-> **NOTE:** Available since v1.37.0.

-> **NOTE:** The endpoint of bssopenapi used only support "business.aliyuncs.com" at present.

-> **NOTE:** From version 1.214.0, if `product_type` is set to `ddoscoo` or `ddoscoo_intl`, the provider `region` should be set to `cn-hangzhou`, and if `product_type` is set to `ddosDip`, the provider `region` should be set to `ap-southeast-1`.

## Example Usage

Basic Usage

<div style="display: block;margin-bottom: 40px;"><div class="oics-button" style="float: right;position: absolute;margin-bottom: 10px;">
  <a href="https://api.aliyun.com/terraform?resource=alicloud_ddoscoo_instance&exampleId=458888fd-0444-fbcf-9a70-b1d21d48f7df523857bd&activeTab=example&spm=docs.r.ddoscoo_instance.0.458888fd04&intl_lang=EN_US" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; max-width: 100%;">
  </a>
</div></div>

```terraform
provider "alicloud" {
  region = "cn-hangzhou"
}

variable "name" {
  default = "tf-example"
}

resource "alicloud_ddoscoo_instance" "default" {
  name              = var.name
  base_bandwidth    = "30"
  bandwidth         = "30"
  service_bandwidth = "100"
  port_count        = "50"
  domain_count      = "50"
  product_type      = "ddoscoo"
  period            = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the instance. This name can have a string of `1` to `64` characters.
* `port_count` - (Required) Port retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
* `domain_count` - (Required) Domain retransmission rule count of the instance. At least 50. Increase 5 per step, such as 55, 60, 65. Only support upgrade.
* `base_bandwidth` - (Optional) Base defend bandwidth of the instance. Valid values: `30`, `60`, `100`, `300`, `400`, `500`, `600`. The unit is Gbps. Only support upgrade. **NOTE:** `base_bandwidth` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
* `bandwidth` - (Optional) Elastic defend bandwidth of the instance. This value must be larger than the base defend bandwidth. Valid values: `30`, `60`, `100`, `300`, `400`, `500`, `600`. The unit is Gbps. Only support upgrade. **NOTE:** `bandwidth` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
* `service_bandwidth` - (Optional) Business bandwidth of the instance. At leaset 100. Increased 100 per step, such as 100, 200, 300. The unit is Mbps. Only support upgrade. **NOTE:** `service_bandwidth` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
* `normal_bandwidth` - (Optional, Available since v1.214.0) The clean bandwidth provided by the instance. **NOTE:** `normal_bandwidth` is valid only when `product_type` is set to `ddosDip`. From version 1.248.0, `normal_bandwidth` can be modified.
* `normal_qps` - (Optional, Available since v1.214.0) The clean QPS provided by the instance. **NOTE:** `normal_qps` is valid only when `product_type` is set to `ddosDip`. From version 1.248.0, `normal_qps` can be modified.
* `edition_sale` - (Optional, ForceNew, Available since v1.212.0) The mitigation plan of the instance. Default value: `coop`. Valid values:
  - `coop`: Anti-DDoS Pro instance of the Profession mitigation plan.
* `product_plan` - (Optional, Available since v1.214.0) The mitigation plan of the instance. Valid values:
  - `0`: The Insurance mitigation plan.
  - `1`: The Unlimited mitigation plan.
  - `2`: The Chinese Mainland Acceleration (CMA) mitigation plan.
  - `3`: The Secure Chinese Mainland Acceleration (Sec-CMA) mitigation plan.
**NOTE:** `product_plan` is valid only when `product_type` is set to `ddosDip`. From version 1.248.0, `product_plan` can be modified.
* `address_type` - (Optional, ForceNew, Available since v1.212.0) The IP version of the IP address. Default value: `Ipv4`. Valid values: `Ipv4`, `Ipv6`. **NOTE:** `address_type` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
* `bandwidth_mode` - (Optional, Available since v1.212.0) The mitigation plan of the instance. Valid values:
  - `0`: Disables the burstable clean bandwidth feature.
  - `1`: Enables the burstable clean bandwidth feature and uses the daily 95th percentile metering method.
  - `2`: Enables the burstable clean bandwidth feature and uses the monthly 95th percentile metering method.
**NOTE:** `bandwidth_mode` is valid only when `product_type` is set to `ddoscoo` or `ddoscoo_intl`.
* `function_version` - (Optional, Available since v1.214.0) The function plan of the instance. Valid values:
  - `0`: The Standard function plan.
  - `1`: The Enhanced function plan.
**NOTE:** From version 1.248.0, `function_version` can be modified.
* `product_type` - (Optional, Available since v1.125.0) The product type for purchasing DDOSCOO instances used to differ different account type. Default value: `ddoscoo`. Valid values:
  - `ddoscoo`: Anti-DDoS Pro. Only supports domestic account.
  - `ddoscoo_intl`: Anti-DDoS Pro. Only supports to international account.
  - `ddosDip`: Anti-DDoS Premium.
**NOTE:** From version 1.214.0, `product_type` can be set to `ddosDip`.
* `period` - (Optional, Int) The duration that you will buy DdosCoo instance (in month). Valid values: [1~9], `12`, `24`, `36`. Default value: `1`. At present, the provider does not support modify `period`.
* `tags` - (Optional, Map, Available since v1.248.0) A mapping of tags to assign to the resource.
* `modify_type` - (Optional, Available since v1.248.0) The type of modification. Valid values: `UPGRADE`, `DOWNGRADE`.

## Attributes Reference

The following attributes are exported:

* `id` - The resource ID in terraform of DdosCoo.
* `ip` - (Available since v1.212.0) The IP address of the Instance.
* `status` - (Available since v1.248.0) The status of the instance.
* `create_time` - (Available since v1.248.0) The time when the instance was created.

## Timeouts

-> **NOTE:** Available since v1.212.0.

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 18 mins) Used when create the Instance.
* `update` - (Defaults to 6 mins) Used when update the Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Instance.

## Import

DdosCoo Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_ddoscoo_instance.example <id>
```
