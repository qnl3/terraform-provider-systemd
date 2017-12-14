package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHostnamed() *schema.Resource {
	return &schema.Resource {
		Create: resourceHostnamedCreate,
		Read: resourceHostnamedRead,
		Update: resourceHostnamedUpdate,
		Delete: resourceHostnamedDelete,

		Schema: map[string]*schema.Schema {
			"user": &schema.Schema {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"hostname": &schema.Schema {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"new_hostname": &schema.Schema {
				Type: schema.TypeString,
				Optional: true,
			},

			"icon_name": &schema.Schema {
				Type: schema.TypeString,
				Optional: true,
			},

			"chassis_name": &schema.Schema {
				Type: schema.TypeString,
				Optional: true,
			},

			"deployment": &schema.Schema {
				Type: schema.TypeString,
				Optional: true,
			},

			"location": &schema.Schema {
				Type: schema.TypeString,
				Optional: true,
			},

		},
	}
}

func resourceHostnamedCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceHostnamedRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceHostnamedUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceHostnamedDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
