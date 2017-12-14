/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAgentGraphDBNodeConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"analytics_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAnalyticsProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"application_persistence_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApplicationPersistenceProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"application_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApplicationProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cloud": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslb": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslbgeodbprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbGeoDbProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslbservice": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbServiceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"health_monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"http_request_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRequestPolicySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"http_response_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPResponsePolicySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"http_security_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPSecurityPolicySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ip_addr_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrGroupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"microservice": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMicroServiceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"network_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNetworkProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"network_security_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNetworkSecurityPolicySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolGroupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"priority_labels": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePriorityLabelsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"serviceenginegroup": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServiceEngineGroupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_key_and_certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLKeyAndCertificateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"string_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceStringGroupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTenantSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"virtual_service_se": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceSeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_data_script": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVSDataScriptSetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
