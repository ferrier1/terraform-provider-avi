/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRPCRequestFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_all_tenant_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAPICEpgAllTenantFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"apic_cli_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApicCliLoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"apic_epg_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAPICEpgFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"arp_table_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceArpTableFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"candidate_se_host_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCandidateSeHostFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteOpsConsistencyCheckerSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"con_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeConsumerFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"connection_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConnectionFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"connpool_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConnpoolFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"core_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCoreFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"corenum_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCorenumFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			//"cps_doser_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceCpsDoserFilterSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			"delete_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDeleteFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			//"flowtable_entry_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceFlowtableEntryFilterSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			//"flowtable_intf_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceFlowtableIntfFilterSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			//"geo_location_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceGeoLocationFilterSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			//"glb_params_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceGlbParamsFilterSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			"gs_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGsFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			//"gs_params_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceGSParamsFilterSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			"httpcache_obj_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHttpCacheObjFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"listeningsock_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceListeningsockFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			//"metrics_agent_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceMetricsAgentFilterSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			//"metrics_mgr_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceMetricsMgrFilterSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			"ms_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMicroServiceFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"nsx_internal_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNsxInternalParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"nsx_sg_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNsxSgFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			//"persistence_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourcePersistenceFilterSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			"placement_status_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePlacementStatusFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"res_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeResourceFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_metrics_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeMetricsFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_params_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeParamsFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"server_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tcp_stat_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTcpStatFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"udp_stat_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceUdpStatFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vcenter_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVcenterLoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vi_datastore_filtler": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVIDatastoreFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vi_host_resource_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVIHostResourceFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vi_network_subnet_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVINetworkSubnetFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vi_redis_datastore_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVIRedisDatastoreFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vi_retrieve_pg_names": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVIRetrievePGNamesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vi_subfolder_filtler": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVISubfolderFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vip_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeVipFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vrf_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePlacementVrfFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_metrics_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVsMetricsFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vstype_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVstypeFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
