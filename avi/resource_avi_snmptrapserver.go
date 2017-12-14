/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSnmpTrapServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"user": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSnmpV3UserParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_VER2",
			},
		},
	}
}
