/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAlertRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"conn_app_log_rule": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAlertFilterSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"event_match_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metrics_rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAlertRuleMetricSchema(),
			},
			"operator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OPERATOR_AND",
			},
			"sys_event_rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAlertRuleEventSchema(),
			},
		},
	}
}
