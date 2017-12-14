/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthScoreResourcesDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_resources": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceApplicationResourcesScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"controller_resources": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceControllerResourcesScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_resources": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolResourcesScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reason_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_resources": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServerResourcesScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"serviceengine_resources": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceServiceEngineResourcesScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
			"virtualservice_resources": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceResourcesScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
