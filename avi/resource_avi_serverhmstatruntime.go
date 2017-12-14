/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServerHMStatRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_addr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"last_transition_timestamp_1": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"last_transition_timestamp_2": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"last_transition_timestamp_3": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeStampSchema(),
				Set: func(v interface{}) int {
					return 0
				},
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
				Optional: true,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shm_runtime": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerHealthMonitorRuntimeSchema(),
			},
		},
	}
}
