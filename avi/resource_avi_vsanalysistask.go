/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsAnalysisTaskSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConfigAnalysisSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"metrics": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsMetricsAnalysisSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"operational": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsOperationalAnalysisSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"placement": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsPlacementAnalysisSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"task_status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"task_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"traffic": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsTrafficAnalysisSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
