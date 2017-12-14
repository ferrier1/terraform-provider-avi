/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIptableRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"dnat_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dst_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dst_port": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"input_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"output_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"src_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"src_port": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortRangeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
