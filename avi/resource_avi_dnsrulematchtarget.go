/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDnsRuleMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip_address": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsClientIpMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"geo_location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsGeoLocationMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsTransportProtocolMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query_name": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsQueryNameMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query_type": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsQueryTypeMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
