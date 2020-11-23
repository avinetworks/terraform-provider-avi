/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviDynamicDnsRecord() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviDynamicDnsRecordRead,
		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cname": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceDnsCnameRdataSchema(),
			},
			"delegated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_vs_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6_address": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsAAAARdataSchema(),
			},
			"ip_address": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsARdataSchema(),
			},
			"metadata": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mx_records": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsMxRdataSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ns": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsNsRdataSchema(),
			},
			"num_records_in_response": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"service_locators": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsSrvRdataSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"txt_records": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsTxtRdataSchema(),
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard_match": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}
