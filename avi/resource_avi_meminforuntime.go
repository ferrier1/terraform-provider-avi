/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMeminfoRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"memory_config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMemConfigSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"process_memory_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMallstatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"protobuf_memory_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMemProtoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shared_memory_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceJestatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
