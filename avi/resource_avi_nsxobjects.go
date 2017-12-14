/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxObjectsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"firewallconfiguration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcensxFwConfigSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ipnodes": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcensxIpNodesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ipset": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceipsetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"list": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcensxObjectListSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pagedsystemeventlist": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcensxSystemEventListSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rule": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceruleSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sections": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcesectionsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"versions": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcensxVersionsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vshieldappconfiguration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcensxVshieldAppConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
