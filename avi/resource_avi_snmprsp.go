/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSnmpRspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_mib_entry": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceControllerMibEntrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rpc_status": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"serviceengine_mib_entry": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServiceEngineMibEntrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"virtualservice_mib_entry": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceMibEntrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
