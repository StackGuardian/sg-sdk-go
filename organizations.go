// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"

	internal "github.com/StackGuardian/sg-sdk-go/internal"
)

type OrgGetResponse struct {
	Msg *Organization `json:"msg,omitempty" url:"msg,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (o *OrgGetResponse) GetMsg() *Organization {
	if o == nil {
		return nil
	}
	return o.Msg
}

func (o *OrgGetResponse) GetExtraProperties() map[string]interface{} {
	return o.extraProperties
}

func (o *OrgGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler OrgGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OrgGetResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	o.rawJSON = json.RawMessage(data)
	return nil
}

func (o *OrgGetResponse) String() string {
	if len(o.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(o.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(o); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", o)
}

type Organization struct {
	ResourceName             *string        `json:"ResourceName,omitempty" url:"ResourceName,omitempty"`
	Admins                   []string       `json:"Admins,omitempty" url:"Admins,omitempty"`
	MarketplaceSubscriptions []interface{}  `json:"MarketplaceSubscriptions,omitempty" url:"MarketplaceSubscriptions,omitempty"`
	IsArchive                *string        `json:"IsArchive,omitempty" url:"IsArchive,omitempty"`
	IsActive                 *IsArchiveEnum `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	ResourceId               *string        `json:"ResourceId,omitempty" url:"ResourceId,omitempty"`
	ModifiedAt               *float64       `json:"ModifiedAt,omitempty" url:"ModifiedAt,omitempty"`
	ParentId                 *string        `json:"ParentId,omitempty" url:"ParentId,omitempty"`
	ResourceType             *string        `json:"ResourceType,omitempty" url:"ResourceType,omitempty"`
	DocVersion               *string        `json:"DocVersion,omitempty" url:"DocVersion,omitempty"`
	Authors                  []string       `json:"Authors,omitempty" url:"Authors,omitempty"`
	ActivitySubscribers      []string       `json:"ActivitySubscribers,omitempty" url:"ActivitySubscribers,omitempty"`
	SubResourceId            *string        `json:"SubResourceId,omitempty" url:"SubResourceId,omitempty"`
	OrgId                    *string        `json:"OrgId,omitempty" url:"OrgId,omitempty"`
	CreatedAt                *float64       `json:"CreatedAt,omitempty" url:"CreatedAt,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (o *Organization) GetResourceName() *string {
	if o == nil {
		return nil
	}
	return o.ResourceName
}

func (o *Organization) GetAdmins() []string {
	if o == nil {
		return nil
	}
	return o.Admins
}

func (o *Organization) GetMarketplaceSubscriptions() []interface{} {
	if o == nil {
		return nil
	}
	return o.MarketplaceSubscriptions
}

func (o *Organization) GetIsArchive() *string {
	if o == nil {
		return nil
	}
	return o.IsArchive
}

func (o *Organization) GetIsActive() *IsArchiveEnum {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *Organization) GetResourceId() *string {
	if o == nil {
		return nil
	}
	return o.ResourceId
}

func (o *Organization) GetModifiedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *Organization) GetParentId() *string {
	if o == nil {
		return nil
	}
	return o.ParentId
}

func (o *Organization) GetResourceType() *string {
	if o == nil {
		return nil
	}
	return o.ResourceType
}

func (o *Organization) GetDocVersion() *string {
	if o == nil {
		return nil
	}
	return o.DocVersion
}

func (o *Organization) GetAuthors() []string {
	if o == nil {
		return nil
	}
	return o.Authors
}

func (o *Organization) GetActivitySubscribers() []string {
	if o == nil {
		return nil
	}
	return o.ActivitySubscribers
}

func (o *Organization) GetSubResourceId() *string {
	if o == nil {
		return nil
	}
	return o.SubResourceId
}

func (o *Organization) GetOrgId() *string {
	if o == nil {
		return nil
	}
	return o.OrgId
}

func (o *Organization) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Organization) GetExtraProperties() map[string]interface{} {
	return o.extraProperties
}

func (o *Organization) UnmarshalJSON(data []byte) error {
	type unmarshaler Organization
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = Organization(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	o.rawJSON = json.RawMessage(data)
	return nil
}

func (o *Organization) String() string {
	if len(o.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(o.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(o); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", o)
}
