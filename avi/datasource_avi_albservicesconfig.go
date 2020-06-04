/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviALBServicesConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviALBServicesConfigRead,
		Schema: map[string]*schema.Schema{
			"asset_contact": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceALBServicesUserSchema(),
			},
			"feature_opt_in_status": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePortalFeatureOptInSchema(),
			},
			"polling_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"portal_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proactive_support_defaults": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceProactiveSupportDefaultsSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
