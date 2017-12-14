/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeResVerifyBindingReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"con_info": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceSeResourceFindReqSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"update_consumer": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
