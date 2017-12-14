/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePlacementOpInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"create_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCreateInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mod_vnic_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceModVnicInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"switchover_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSwitchoverInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip_vnic_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceUpdVipVnicInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vnic_ip_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVnicIpInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
