/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRPCRequestResourceSchema() *schema.Resource {
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
			"apic_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAPICConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"apic_transaction": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAPICTransactionSchema(),
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
			"auth_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAuthProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"auto_scale_launch_config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAutoScaleLaunchConfigSchema(),
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
			"cloud_props": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudPropertiesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"controller_props": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceControllerPropertiesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"debug_controller": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDebugControllerSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"debug_service_engine": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDebugServiceEngineSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dns_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsPolicySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"errorpage_body": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceErrorPageBodySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"errorpage_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceErrorPageProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			//"gs_ops": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceGslbSiteOpsSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			"gslb_dns_geo_update": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDnsGeoUpdateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslb_dns_update": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbDnsUpdateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslb_geodb": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbGeoDbProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslb_service": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbServiceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslb_service_runtime": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbServiceRuntimeSchema(),
			},
			"hardwaresecuritymodule_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHardwareSecurityModuleGroupSchema(),
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
			"http_policy_set": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPPolicySetSchema(),
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
			"ipam_dns_records": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsRecordSchema(),
			},
			"microservice": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMicroServiceSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"microservice_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMicroServiceGroupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"network": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNetworkSchema(),
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
			"pki_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePKIProfileSchema(),
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
			"pool_group_deployment_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolGroupDeploymentPolicySchema(),
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
			"se_props": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSePropertiesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_auto_scale_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerAutoScalePolicySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"service_engine": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServiceEngineSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"service_engine_group": &schema.Schema{
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
			"system_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSystemConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"traffic_clone_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTrafficCloneProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vrf_context": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVrfContextSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vsdatascriptset": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVSDataScriptSetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"waf_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWafPolicySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"waf_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWafProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
