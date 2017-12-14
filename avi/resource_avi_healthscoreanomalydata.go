/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthScoreAnomalyDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_anomaly": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApplicationAnomalyScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"controller_anomaly": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceControllerAnomalyScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_anomaly": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolAnomalyScoreDataSchema(),
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
			"server_anomaly": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAnomalyScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"serviceengine_anomaly": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServiceEngineAnomalyScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
			"virtualservice_anomaly": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceAnomalyScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
