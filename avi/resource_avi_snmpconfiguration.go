/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSnmpConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"snmp_v3_config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSnmpV3ConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sys_contact": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "support@avinetworks.com",
			},
			"sys_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_VER2",
			},
		},
	}
}
