/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxSgDynamicMemberSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dynamicset": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourcensxDynamicSetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
