/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNetworkSecurityRuleDosSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"match": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Elem: ResourceNetworkSecurityMatchTargetSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rl_param": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNetworkSecurityPolicyActionRLParamSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
