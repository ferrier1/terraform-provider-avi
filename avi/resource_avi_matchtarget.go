/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMatchTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cookie": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCookieMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hdrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHdrMatchSchema(),
			},
			"host_hdr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHostHdrMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"method": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMethodMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"path": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePathMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProtocolMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceQueryMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"version": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPVersionMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vs_port": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePortMatchSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
