/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMesosConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_vses_are_feproxy": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"app_sync_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"container_port_match_http_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"coredump_directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump",
			},
			"disable_auto_backend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_frontend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_gs_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_se_creation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_registry_se": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDockerRegistrySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"east_west_placement_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"enable_event_subscription": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_bridge_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cbr1",
			},
			"feproxy_container_port_as_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_route_publish": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFeProxyRoutePublishConfigSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"feproxy_vips_enable_proxy_arp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fleet_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_container_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"marathon_configurations": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMarathonConfigurationSchema(),
			},
			"marathon_se_deployment": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceMarathonSeDeploymentSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"mesos_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http://leader.mesos:5050",
			},
			"node_availability_zone_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_controller": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceNuageSDNControllerSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"se_deployment_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "MESOS_SE_CREATE_FLEET",
			},
			"se_exclude_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_include_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema(),
			},
			"se_resources": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosSeResourcesSchema(),
			},
			"se_spawn_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"se_volume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/opt/avi/se",
			},
			"services_accessible_all_interfaces": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssh_user_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_bridge_ip_as_vip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_container_ip_port": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_controller_image": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_vips_for_east_west_services": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
