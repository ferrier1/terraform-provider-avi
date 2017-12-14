/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAlgoStatRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ca_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCoreAffinityStatSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"lb_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lc_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLeastConnectionStatSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ns_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNsStatSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"wrr_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWrrStatSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
