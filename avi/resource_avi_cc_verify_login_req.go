/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_verify_login_reqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aws_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAWSLoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"azure_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudConnectorUserSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cs_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCloudStackLoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"docker_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDockerLoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ms_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMesosLoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"os_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOpenstackLoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"oshiftk8s_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOShiftK8SLoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"rancher_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRancherLoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vca_login": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceVCALoginSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
