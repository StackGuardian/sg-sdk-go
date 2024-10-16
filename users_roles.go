// This file was auto-generated by Fern from our API Definition.

package api

type PatchedRole struct {
	ResourceName       *string                        `json:"ResourceName,omitempty" url:"-"`
	ParentId           *string                        `json:"ParentId,omitempty" url:"-"`
	ResourceId         *string                        `json:"ResourceId,omitempty" url:"-"`
	ResourceType       *string                        `json:"ResourceType,omitempty" url:"-"`
	Authors            []interface{}                  `json:"Authors,omitempty" url:"-"`
	Description        *string                        `json:"Description,omitempty" url:"-"`
	Tags               []string                       `json:"Tags,omitempty" url:"-"`
	AllowedPermissions map[string]*AllowedPermissions `json:"AllowedPermissions,omitempty" url:"-"`
}
