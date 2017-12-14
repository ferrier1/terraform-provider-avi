/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_AgentInternalSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"async": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_AsyncSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"aws": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_AWSSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"azure": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_AzureSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"bgworkers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_BGWorkerSchema(),
			},
			"cc_flags": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_AgentFlagsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cc_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cc_threadid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cc_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloudstack": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_CloudStackSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"config_genid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovery": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_DiscoverySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"docker": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_DockerSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"format": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_FormatSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"health": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_HealthSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"images": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_ImageSchema(),
			},
			"init": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_InitSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ipencap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"linuxserver": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_LinuxServerSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mesos": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_MesosSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"openstack": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_OpenStackSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"oshift_k8s": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_OShiftK8SSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_PortsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"prefer_avi_rt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProxyConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"proxy_provider": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rancher": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_RancherSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ret_status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ret_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"up_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcloudair": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_VCloudAirSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
