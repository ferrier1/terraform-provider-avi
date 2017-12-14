/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAnomalyContextSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ema_context": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceExponentialMovingAverageCtxSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ewma_context": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceExponentialMovingAverageCtxSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hw_at_as_context": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHoltWintersCtxSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hw_at_ms_context": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHoltWintersCtxSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"prediction_interval_high": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "0.0",
			},
			"prediction_interval_low": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "0.0",
			},
		},
	}
}
