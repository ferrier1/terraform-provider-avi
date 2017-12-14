/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceTaskBaseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"db_notification": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTaskDbNotificationBaseSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"request": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTaskRequestBaseSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"response": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTaskResponseBaseSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sender_node_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
