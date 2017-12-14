/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeConsumerResyncRespSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"res_rsp": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeResourceFindRespSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}
