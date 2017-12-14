/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRPCRequestActionParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apic_disable_enable_vs_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAPICDisableEnableVsParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"apic_txn_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAPICTransactionFlapParamsSchema(),
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
			"connection_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConnectionClearFilterSchema(),
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
			//"cps_doser_filter": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceCpsDoserFilterSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			//"dispatcher_stat_clear": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceDispatcherStatClearSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			//"dispatcher_table_dump_clear": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceDispatcherTableDumpClearSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			"event_gen_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceEventGenParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			//"ft_entry_delete": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceFlowtableEntryFilterSchema(),
			//	Set: func(v interface{}) int {
			//		return 0
			//	},
			//},
			//"gs_ops": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem:     ResourceGslbSiteOpsSchema(),
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
			"nsx_fault_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNsxFaultParamsSchema(),
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
			"pool_server_state": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolServerStateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rediscover_vcenter_param": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRediscoverVcenterParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"resync": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbSiteOpsResyncSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"retry_placement_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRetryPlacementParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_fault_inject_exhaust_param": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSEFaultInjectExhaustParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_switchover_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeSwitchoverParamsSchema(),
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
			"vi_delete_network_filter": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVIDeleteNetworkFilterSchema(),
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
		},
	}
}
