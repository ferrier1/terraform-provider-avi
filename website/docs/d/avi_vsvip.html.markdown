---
layout: "avi"
page_title: "AVI: avi_vsvip"
sidebar_current: "docs-avi-datasource-vsvip"
description: |-
  Get information of Avi VsVip.
---

# avi_vsvip

This data source is used to to get avi_vsvip objects.

## Example Usage

```hcl
data "VsVip" "foo_VsVip" {
    uuid = "VsVip-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search VsVip by name.
* `uuid` - (Optional) Search VsVip by uuid.
* `cloud_ref` - (Optional) Search VsVip by cloud_ref.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_ref` - It is a reference to an object of type cloud.
* `dns_info` - Service discovery specific data including fully qualified domain name, type and time-to-live of the dns record.
* `east_west_placement` - Force placement on all service engines in the service engine group (container clouds only).
* `name` - Name for the vsvip object.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the vsvip object.
* `vip` - List of virtual service ips and other shareable entities.
* `vrf_context_ref` - Virtual routing context that the virtual service is bound to.
* `vsvip_cloud_config_cksum` - Checksum of cloud configuration for vsvip.
