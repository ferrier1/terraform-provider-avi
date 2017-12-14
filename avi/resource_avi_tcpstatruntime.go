/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceTcpStatRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connection_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConnectionStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"connections_dropped": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceConnectionDropStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dns_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTcpDnsStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ecn_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceEcnStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"misc_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMiscStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"packets_dropped": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePacketDropStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"retransmit_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRetransmitStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rx_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRxStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sack_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSackStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"syncache_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSyncacheStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTimeoutStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tx_stats": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTxStatsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
