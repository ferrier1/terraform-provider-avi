/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNetworkManagerDiagRspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"networks": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVcenterNetworkDiagSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"success": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"summary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Success",
			},
		},
	}
}
