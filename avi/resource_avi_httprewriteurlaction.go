/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHTTPRewriteURLActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_hdr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"path": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"query": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamQuerySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
