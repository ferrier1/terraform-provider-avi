/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIcmpStatRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"icps_inhist": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIcmpHistSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"icps_outhist": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIcmpHistSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"more_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIcmpStatSchema(),
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
		},
	}
}
