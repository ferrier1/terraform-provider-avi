/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeAgentGraphDBRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"analyticsprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"applicationpersistenceprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"applicationprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cloud": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"graph_version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"gslb": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslbgeodbprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gslbservice": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"healthmonitor": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"httprequestpolicy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"httpresponsepolicy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"httpsecuritypolicy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ipaddrgroup": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"microservice": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"networkprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"networksecuritypolicy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"serviceenginegroup": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sslkeyandcertificate": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"sslprofile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"stringgroup": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tenant": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"total_obj": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"total_obj_active": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"total_obj_awaiting_dp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"total_obj_error": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"total_obj_ew_subnet_error": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"virtualservice": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"vsdatascriptset": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeAgentGraphDBNodeInfoSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
