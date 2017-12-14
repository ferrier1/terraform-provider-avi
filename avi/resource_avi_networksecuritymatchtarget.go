/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNetworkSecurityMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"microservice": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMicroServiceMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_port": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
