/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricObjTypesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anomaly": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsAnomalyObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"app_insights": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsAppInsightsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"collection": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricCollectionsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"controller_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceControllerStatsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dns_client": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsVserverDNSObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dns_server": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsServerDNSObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dos": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsDosAnalyticsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"healthscore": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsHealthScoreObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"l4_client": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVserverL4MetricsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"l4_server": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerL4MetricsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"l7_client": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVserverL7MetricsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"l7_server": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerL7MetricsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"process_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProcessStatsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rum": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsRumAnalyticsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rum_browser": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsRumPreaggBrowserObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rum_country": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsRumPreaggCountryObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rum_devtype": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsRumPreaggDevtypeObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rum_ipgroup": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsRumPreaggIPGroupObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rum_lang": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsRumPreaggLangObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rum_os": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsRumPreaggOsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rum_url": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsRumPreaggUrlObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"samples_blob": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsRumSamplesBlobObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"samples_dimension": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsRumSamplesDimensionObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_if": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsSeIfStatsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeStatsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"service_insights": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsServiceInsightsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"source_insights": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsSourceInsightsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tenant_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsTenantStatsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"user_metrics": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsUserMetricsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vm_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVmStatsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"waf_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsWafGroupMetricsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"waf_rule": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsWafRuleMetricsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"waf_tag": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMetricsWafTagMetricsObjSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
