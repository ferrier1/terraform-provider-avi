/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDebugFilterUnionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alert_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAlertMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"autoscale_mgr_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAutoScaleMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cloud_connector_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudConnectorDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hs_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHSMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mesos_metrics_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMesosMetricsDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"metrics_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_mgr_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"state_cache_mgr_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStateCacheMgrDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vs_debug_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsDebugFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
