// This file was auto-generated by Fern from our API Definition.

package api

type PatchedPolicy struct {
	// Resource Name of the policy
	ResourceName *string `json:"ResourceName,omitempty" url:"-"`
	// Description of the policy
	Description *string `json:"Description,omitempty" url:"-"`
	// List of IDs of the approvers for the policy
	Approvers []string `json:"Approvers,omitempty" url:"-"`
	// Number of approvals required for the policy to be enforced
	NumberOfApprovalsRequired *int `json:"NumberOfApprovalsRequired,omitempty" url:"-"`
	// Tags for the policy
	Tags []string `json:"Tags,omitempty" url:"-"`
	// Should the policy be active?, choices are 0 or 1
	//
	// * `0` - 0
	// * `1` - 1
	IsActive *IsActiveEnum `json:"IsActive,omitempty" url:"-"`
	// What the policy will be enforced on.
	EnforcedOn []string `json:"EnforcedOn,omitempty" url:"-"`
	// Policies Config for the policy
	PoliciesConfig []*PoliciesConfig `json:"PoliciesConfig,omitempty" url:"-"`
}
