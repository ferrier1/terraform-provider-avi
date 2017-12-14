/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClusterLeaderFailoverEventSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"leader_node": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"previous_leader_node": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceClusterNodeSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
