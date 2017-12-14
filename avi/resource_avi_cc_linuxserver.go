/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_LinuxServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cfg": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLinuxServerConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_HostSchema(),
			},
			"inited_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ipam_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_IpamInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sshinfo": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_SshInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"uninited_hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"version_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
