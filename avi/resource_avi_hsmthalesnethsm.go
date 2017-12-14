/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHSMThalesNetHsmSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"esn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"keyhash": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"module_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"remote_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  9004,
			},
		},
	}
}
