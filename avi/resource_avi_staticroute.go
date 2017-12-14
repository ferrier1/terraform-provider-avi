/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceStaticRouteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_gateway_monitor": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"next_hop": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"route_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
