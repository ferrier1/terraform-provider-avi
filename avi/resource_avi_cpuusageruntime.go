/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCpuUsageRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"current": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceCpuUsageEntrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"last_5_sec_period": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceCpuUsageEntrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"since_boot": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceCpuUsageEntrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
