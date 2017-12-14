/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbSiteRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rxed_site_hs": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteHealthStatusSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"site_cfg": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeCfgSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"site_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"site_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteRuntimeStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
