/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNsxFaultSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"del_fault": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceRestApiFaultSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"get_fault": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceRestApiFaultSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"post_fault": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceRestApiFaultSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"put_fault": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceRestApiFaultSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
