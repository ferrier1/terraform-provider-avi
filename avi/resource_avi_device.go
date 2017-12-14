/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDeviceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCpuSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"network": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVmNetworkSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"storage": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStorageSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
