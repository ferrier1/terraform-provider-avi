/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceConnpoolServerInternalSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"conn_bind": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceConnectionItemSchema(),
			},
			"conn_free": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceConnectionItemSchema(),
			},
			"ip_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"max_connection": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_conn_bind": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_conn_free": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOperationalStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConnpoolStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
