/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVirtualServiceResourcesScoreSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hs_entity": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceHealthScoreEntitySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"score_data": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVirtualServiceResourcesScoreDataSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
		},
	}
}
