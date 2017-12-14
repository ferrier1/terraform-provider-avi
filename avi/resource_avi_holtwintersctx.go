/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHoltWintersCtxSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deviation": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "0.0",
			},
			"level": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "0.0",
			},
			"prediction": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "0.0",
			},
			"prediction_interval_high": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "0.0",
			},
			"prediction_interval_low": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "0.0",
			},
			"trend": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  "0.0",
			},
		},
	}
}
