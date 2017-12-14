/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_update_vip_reqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cloud-0",
			},
			"free_old": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"new": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     Resourcecc_vip_infoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"old": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     Resourcecc_vip_infoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
