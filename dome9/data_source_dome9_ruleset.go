package dome9

import (
	log "github.com/sourcegraph-ce/logrus"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRuleSet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRuleSetRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"language": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hide_in_compliance": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_template": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"min_feature_tier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logic": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"severity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remediation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compliance_tag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"control_title": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rule_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"logic_hash": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRuleSetRead(d *schema.ResourceData, meta interface{}) error {
	d9Client := meta.(*Client)

	id := d.Get("id").(string)
	log.Printf("[INFO] Getting data for rule set id %s\n", id)

	resp, _, err := d9Client.ruleSet.Get(id)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(resp.ID))
	_ = d.Set("name", resp.Name)
	_ = d.Set("description", resp.Description)
	_ = d.Set("cloud_vendor", resp.CloudVendor)
	_ = d.Set("language", resp.Language)
	_ = d.Set("hide_in_compliance", resp.HideInCompliance)
	_ = d.Set("is_template", resp.IsTemplate)
	_ = d.Set("min_feature_tier", resp.MinFeatureTier)
	_ = d.Set("created_time", resp.CreatedTime)
	_ = d.Set("updated_time", resp.UpdatedTime)

	if err := d.Set("rules", flattenRules(resp.Rules)); err != nil {
		return err
	}

	return nil
}
