// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"

	core "github.com/StackGuardian/sg-sdk-go/core"
	internal "github.com/StackGuardian/sg-sdk-go/internal"
)

type Role struct {
	ResourceName string                   `json:"ResourceName" url:"-"`
	Description  *core.Optional[string]   `json:"Description,omitempty" url:"-"`
	Tags         *core.Optional[[]string] `json:"Tags,omitempty" url:"-"`
	// Contextual tags to give context to your tags
	ContextTags        *core.Optional[map[string]*string]             `json:"ContextTags,omitempty" url:"-"`
	AllowedPermissions *core.Optional[map[string]*AllowedPermissions] `json:"AllowedPermissions,omitempty" url:"-"`
	DocVersion         *core.Optional[DocVersionEnum]                 `json:"DocVersion,omitempty" url:"-"`
}

type PatchedRole struct {
	ResourceName *core.Optional[string]   `json:"ResourceName,omitempty" url:"-"`
	Description  *core.Optional[string]   `json:"Description,omitempty" url:"-"`
	Tags         *core.Optional[[]string] `json:"Tags,omitempty" url:"-"`
	// Contextual tags to give context to your tags
	ContextTags        *core.Optional[map[string]*string]             `json:"ContextTags,omitempty" url:"-"`
	AllowedPermissions *core.Optional[map[string]*AllowedPermissions] `json:"AllowedPermissions,omitempty" url:"-"`
	DocVersion         *core.Optional[DocVersionEnum]                 `json:"DocVersion,omitempty" url:"-"`
}

type AddUserToOrganization struct {
	UserId       string          `json:"userId" url:"userId"`
	EntityType   *EntityTypeEnum `json:"entityType,omitempty" url:"entityType,omitempty"`
	Role         *string         `json:"role,omitempty" url:"role,omitempty"`
	Roles        []string        `json:"roles,omitempty" url:"roles,omitempty"`
	ResendInvite *bool           `json:"resendInvite,omitempty" url:"resendInvite,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AddUserToOrganization) GetUserId() string {
	if a == nil {
		return ""
	}
	return a.UserId
}

func (a *AddUserToOrganization) GetEntityType() *EntityTypeEnum {
	if a == nil {
		return nil
	}
	return a.EntityType
}

func (a *AddUserToOrganization) GetRole() *string {
	if a == nil {
		return nil
	}
	return a.Role
}

func (a *AddUserToOrganization) GetRoles() []string {
	if a == nil {
		return nil
	}
	return a.Roles
}

func (a *AddUserToOrganization) GetResendInvite() *bool {
	if a == nil {
		return nil
	}
	return a.ResendInvite
}

func (a *AddUserToOrganization) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AddUserToOrganization) UnmarshalJSON(data []byte) error {
	type unmarshaler AddUserToOrganization
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AddUserToOrganization(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AddUserToOrganization) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AddUserToOrganizationResponse struct {
	Msg  *string                `json:"msg,omitempty" url:"msg,omitempty"`
	Data *AddUserToOrganization `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AddUserToOrganizationResponse) GetMsg() *string {
	if a == nil {
		return nil
	}
	return a.Msg
}

func (a *AddUserToOrganizationResponse) GetData() *AddUserToOrganization {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *AddUserToOrganizationResponse) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AddUserToOrganizationResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AddUserToOrganizationResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AddUserToOrganizationResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AddUserToOrganizationResponse) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type AllowedPermissions struct {
	Name  string              `json:"name" url:"name"`
	Paths map[string][]string `json:"paths,omitempty" url:"paths,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AllowedPermissions) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

func (a *AllowedPermissions) GetPaths() map[string][]string {
	if a == nil {
		return nil
	}
	return a.Paths
}

func (a *AllowedPermissions) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AllowedPermissions) UnmarshalJSON(data []byte) error {
	type unmarshaler AllowedPermissions
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AllowedPermissions(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AllowedPermissions) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

// * `V3.BETA` - V3.BETA
// * `V4` - V4
type DocVersionEnum string

const (
	DocVersionEnumV3Beta DocVersionEnum = "V3.BETA"
	DocVersionEnumV4     DocVersionEnum = "V4"
)

func NewDocVersionEnumFromString(s string) (DocVersionEnum, error) {
	switch s {
	case "V3.BETA":
		return DocVersionEnumV3Beta, nil
	case "V4":
		return DocVersionEnumV4, nil
	}
	var t DocVersionEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DocVersionEnum) Ptr() *DocVersionEnum {
	return &d
}

// * `EMAIL` - EMAIL
// * `GROUP` - GROUP
type EntityTypeEnum string

const (
	EntityTypeEnumEmail EntityTypeEnum = "EMAIL"
	EntityTypeEnumGroup EntityTypeEnum = "GROUP"
)

func NewEntityTypeEnumFromString(s string) (EntityTypeEnum, error) {
	switch s {
	case "EMAIL":
		return EntityTypeEnumEmail, nil
	case "GROUP":
		return EntityTypeEnumGroup, nil
	}
	var t EntityTypeEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EntityTypeEnum) Ptr() *EntityTypeEnum {
	return &e
}

type GetorRemoveUserFromOrganization struct {
	UserId string `json:"userId" url:"userId"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetorRemoveUserFromOrganization) GetUserId() string {
	if g == nil {
		return ""
	}
	return g.UserId
}

func (g *GetorRemoveUserFromOrganization) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetorRemoveUserFromOrganization) UnmarshalJSON(data []byte) error {
	type unmarshaler GetorRemoveUserFromOrganization
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetorRemoveUserFromOrganization(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetorRemoveUserFromOrganization) String() string {
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type RemoveUserFromOrganizationResponse struct {
	Msg  *string                `json:"msg,omitempty" url:"msg,omitempty"`
	Data *AddUserToOrganization `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RemoveUserFromOrganizationResponse) GetMsg() *string {
	if r == nil {
		return nil
	}
	return r.Msg
}

func (r *RemoveUserFromOrganizationResponse) GetData() *AddUserToOrganization {
	if r == nil {
		return nil
	}
	return r.Data
}

func (r *RemoveUserFromOrganizationResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RemoveUserFromOrganizationResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RemoveUserFromOrganizationResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RemoveUserFromOrganizationResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RemoveUserFromOrganizationResponse) String() string {
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

type RoleCreateUpdateResponse struct {
	Msg  *string           `json:"msg,omitempty" url:"msg,omitempty"`
	Data *RoleDataResponse `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RoleCreateUpdateResponse) GetMsg() *string {
	if r == nil {
		return nil
	}
	return r.Msg
}

func (r *RoleCreateUpdateResponse) GetData() *RoleDataResponse {
	if r == nil {
		return nil
	}
	return r.Data
}

func (r *RoleCreateUpdateResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RoleCreateUpdateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RoleCreateUpdateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RoleCreateUpdateResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RoleCreateUpdateResponse) String() string {
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

type RoleDataResponse struct {
	ResourceName string   `json:"ResourceName" url:"ResourceName"`
	Description  *string  `json:"Description,omitempty" url:"Description,omitempty"`
	Tags         []string `json:"Tags,omitempty" url:"Tags,omitempty"`
	// Contextual tags to give context to your tags
	ContextTags        map[string]*string             `json:"ContextTags,omitempty" url:"ContextTags,omitempty"`
	AllowedPermissions map[string]*AllowedPermissions `json:"AllowedPermissions,omitempty" url:"AllowedPermissions,omitempty"`
	DocVersion         string                         `json:"DocVersion" url:"DocVersion"`
	ParentId           string                         `json:"ParentId" url:"ParentId"`
	ResourceId         string                         `json:"ResourceId" url:"ResourceId"`
	ResourceType       string                         `json:"ResourceType" url:"ResourceType"`
	Authors            []interface{}                  `json:"Authors,omitempty" url:"Authors,omitempty"`
	// Time in milliseconds as a string
	CreatedAt int `json:"CreatedAt" url:"CreatedAt"`
	// Time in milliseconds as a string
	ModifiedAt int            `json:"ModifiedAt" url:"ModifiedAt"`
	IsActive   *IsArchiveEnum `json:"IsActive,omitempty" url:"IsActive,omitempty"`
	IsArchive  IsArchiveEnum  `json:"IsArchive" url:"IsArchive"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RoleDataResponse) GetResourceName() string {
	if r == nil {
		return ""
	}
	return r.ResourceName
}

func (r *RoleDataResponse) GetDescription() *string {
	if r == nil {
		return nil
	}
	return r.Description
}

func (r *RoleDataResponse) GetTags() []string {
	if r == nil {
		return nil
	}
	return r.Tags
}

func (r *RoleDataResponse) GetContextTags() map[string]*string {
	if r == nil {
		return nil
	}
	return r.ContextTags
}

func (r *RoleDataResponse) GetAllowedPermissions() map[string]*AllowedPermissions {
	if r == nil {
		return nil
	}
	return r.AllowedPermissions
}

func (r *RoleDataResponse) GetDocVersion() string {
	if r == nil {
		return ""
	}
	return r.DocVersion
}

func (r *RoleDataResponse) GetParentId() string {
	if r == nil {
		return ""
	}
	return r.ParentId
}

func (r *RoleDataResponse) GetResourceId() string {
	if r == nil {
		return ""
	}
	return r.ResourceId
}

func (r *RoleDataResponse) GetResourceType() string {
	if r == nil {
		return ""
	}
	return r.ResourceType
}

func (r *RoleDataResponse) GetAuthors() []interface{} {
	if r == nil {
		return nil
	}
	return r.Authors
}

func (r *RoleDataResponse) GetCreatedAt() int {
	if r == nil {
		return 0
	}
	return r.CreatedAt
}

func (r *RoleDataResponse) GetModifiedAt() int {
	if r == nil {
		return 0
	}
	return r.ModifiedAt
}

func (r *RoleDataResponse) GetIsActive() *IsArchiveEnum {
	if r == nil {
		return nil
	}
	return r.IsActive
}

func (r *RoleDataResponse) GetIsArchive() IsArchiveEnum {
	if r == nil {
		return ""
	}
	return r.IsArchive
}

func (r *RoleDataResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RoleDataResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RoleDataResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RoleDataResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RoleDataResponse) String() string {
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

type RoleGetResponse struct {
	Msg  *RoleDataResponse `json:"msg,omitempty" url:"msg,omitempty"`
	Data *RoleDataResponse `json:"data,omitempty" url:"data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RoleGetResponse) GetMsg() *RoleDataResponse {
	if r == nil {
		return nil
	}
	return r.Msg
}

func (r *RoleGetResponse) GetData() *RoleDataResponse {
	if r == nil {
		return nil
	}
	return r.Data
}

func (r *RoleGetResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RoleGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RoleGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RoleGetResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RoleGetResponse) String() string {
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}
