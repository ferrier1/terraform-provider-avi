/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSystemUpgradeStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller_state": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceControllerUpgradeStateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"from_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"in_progress": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rollback": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_state": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeUpgradeStatusSummarySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"to_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
