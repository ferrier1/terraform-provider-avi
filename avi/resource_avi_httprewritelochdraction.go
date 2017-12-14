/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHTTPRewriteLocHdrActionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"keep_query": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceURIParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
