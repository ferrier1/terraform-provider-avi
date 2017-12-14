/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthScorePerformanceDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_performance": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApplicationPerformanceScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_performance": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolPerformanceScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_performance": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerPerformanceScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"serviceengine_performance": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServiceEnginePerformanceScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
			"virtualservice_performance": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServicePerformanceScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
