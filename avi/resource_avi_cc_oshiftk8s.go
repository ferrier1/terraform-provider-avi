/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_OShiftK8SSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cfg": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOShiftK8SConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"dns_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_DnsInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_HostSchema(),
			},
			"ipam_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_IpamInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"num_pending_apis": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_pending_changes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sshinfo": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_SshInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
