package sdk

import (
	"encoding/json"
	"time"
)

type IdentityProviderPolicy struct {
	Embedded        interface{}           `json:"_embedded,omitempty"`
	Links           interface{}           `json:"_links,omitempty"`
	Conditions      *PolicyRuleConditions `json:"conditions,omitempty"`
	Created         *time.Time            `json:"created,omitempty"`
	Description     string                `json:"description,omitempty"`
	Id              string                `json:"id,omitempty"`
	LastUpdated     *time.Time            `json:"lastUpdated,omitempty"`
	Name            string                `json:"name,omitempty"`
	Priority        int64                 `json:"-"`
	PriorityPtr     *int64                `json:"priority,omitempty"`
	Status          string                `json:"status,omitempty"`
	System          *bool                 `json:"system,omitempty"`
	Type            string                `json:"type,omitempty"`
	AccountLink     *PolicyAccountLink    `json:"accountLink,omitempty"`
	MaxClockSkew    int64                 `json:"-"`
	MaxClockSkewPtr *int64                `json:"maxClockSkew,omitempty"`
	Provisioning    *Provisioning         `json:"provisioning,omitempty"`
	Subject         *PolicySubject        `json:"subject,omitempty"`
}

func NewIdentityProviderPolicy() *IdentityProviderPolicy {
	return &IdentityProviderPolicy{
		Type: "IDP_DISCOVERY",
	}
}

func (a *IdentityProviderPolicy) IsPolicyInstance() bool {
	return true
}

func (a *IdentityProviderPolicy) MarshalJSON() ([]byte, error) {
	type Alias IdentityProviderPolicy
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.Priority != 0 {
		result.PriorityPtr = Int64Ptr(a.Priority)
	}
	if a.MaxClockSkew != 0 {
		result.MaxClockSkewPtr = Int64Ptr(a.MaxClockSkew)
	}
	return json.Marshal(&result)
}

func (a *IdentityProviderPolicy) UnmarshalJSON(data []byte) error {
	type Alias IdentityProviderPolicy

	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	if result.PriorityPtr != nil {
		a.Priority = *result.PriorityPtr
		a.PriorityPtr = result.PriorityPtr
	}
	if result.MaxClockSkewPtr != nil {
		a.MaxClockSkew = *result.MaxClockSkewPtr
		a.MaxClockSkewPtr = result.MaxClockSkewPtr
	}
	return nil
}
