/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbSiteRuntimeCfgSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fd_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gap_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"geo_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ghm_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"glb_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gpki_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gs_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mm_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sync_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteCfgSyncInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
