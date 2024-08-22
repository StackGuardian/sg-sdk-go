// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
)

type CreateTemplateOrRevisionRequest struct {
	// Organization name of the template owner, e.g. my-sg-org
	SgOrgid string    `json:"-" url:"-"`
	Body    *Template `json:"-" url:"-"`
}

func (c *CreateTemplateOrRevisionRequest) UnmarshalJSON(data []byte) error {
	body := new(Template)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	c.Body = body
	return nil
}

func (c *CreateTemplateOrRevisionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Body)
}

type PatchedTemplateUpdate struct {
	// Organization name of the template owner, e.g. my-sg-org
	SgOrgid               string                 `json:"-" url:"-"`
	IsPublic              *IsArchiveEnum         `json:"IsPublic,omitempty" url:"-"`
	LongDescription       *string                `json:"LongDescription,omitempty" url:"-"`
	ShortDescription      *string                `json:"ShortDescription,omitempty" url:"-"`
	Deprecation           *Deprecation           `json:"Deprecation,omitempty" url:"-"`
	SharedOrgsList        []string               `json:"SharedOrgsList,omitempty" url:"-"`
	InputSchemas          []*InputSchemas        `json:"InputSchemas,omitempty" url:"-"`
	Templates             []*Templates           `json:"Templates,omitempty" url:"-"`
	Tags                  []string               `json:"Tags,omitempty" url:"-"`
	GitHubComSync         map[string]interface{} `json:"GitHubComSync,omitempty" url:"-"`
	VcsTriggers           *VcsTriggers           `json:"VCSTriggers,omitempty" url:"-"`
	TerraformIntelligence map[string]interface{} `json:"TerraformIntelligence,omitempty" url:"-"`
	RuntimeSource         *RuntimeSource         `json:"RuntimeSource,omitempty" url:"-"`
}
