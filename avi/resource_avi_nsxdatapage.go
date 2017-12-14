/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxDataPageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"paginginfo": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcensxPagingInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"systemevent": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcensxSystemEventSchema(),
			},
		},
	}
}
