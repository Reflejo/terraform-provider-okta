package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

type Policies interface {
	IsPolicyInstance() bool
}

type PolicyResource resource

type Policy struct {
	Embedded    interface{}           `json:"_embedded,omitempty"`
	Links       interface{}           `json:"_links,omitempty"`
	Conditions  *PolicyRuleConditions `json:"conditions,omitempty"`
	Created     *time.Time            `json:"created,omitempty"`
	Description string                `json:"description,omitempty"`
	Id          string                `json:"id,omitempty"`
	LastUpdated *time.Time            `json:"lastUpdated,omitempty"`
	Name        string                `json:"name,omitempty"`
	Priority    int64                 `json:"-"`
	PriorityPtr *int64                `json:"priority,omitempty"`
	Status      string                `json:"status,omitempty"`
	System      *bool                 `json:"system,omitempty"`
	Type        string                `json:"type,omitempty"`
}

func NewPolicy() *Policy {
	return &Policy{}
}

func (a *Policy) IsPolicyInstance() bool {
	return true
}

// Gets a policy.
func (m *PolicyResource) GetPolicy(ctx context.Context, policyId string, policyInstance Policies, qp *query.Params) (Policies, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v", policyId)
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := policyInstance

	resp, err := rq.Do(ctx, req, &policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// Updates a policy.
func (m *PolicyResource) UpdatePolicy(ctx context.Context, policyId string, body Policies) (Policies, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v", policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	policy := body

	resp, err := rq.Do(ctx, req, &policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// Removes a policy.
func (m *PolicyResource) DeletePolicy(ctx context.Context, policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v", policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Gets all policies with the specified type.
func (m *PolicyResource) ListPolicies(ctx context.Context, qp *query.Params) ([]Policies, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies")
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var policy []Policy

	resp, err := rq.Do(ctx, req, &policy)
	if err != nil {
		return nil, resp, err
	}

	policies := make([]Policies, len(policy))
	for i := range policy {
		policies[i] = &policy[i]
	}
	return policies, resp, nil
}

// Creates a policy.
func (m *PolicyResource) CreatePolicy(ctx context.Context, body Policies, qp *query.Params) (Policies, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies")
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	policy := body

	resp, err := rq.Do(ctx, req, &policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// Activates a policy.
func (m *PolicyResource) ActivatePolicy(ctx context.Context, policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/lifecycle/activate", policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deactivates a policy.
func (m *PolicyResource) DeactivatePolicy(ctx context.Context, policyId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/lifecycle/deactivate", policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Enumerates all policy rules.
func (m *PolicyResource) ListPolicyRules(ctx context.Context, policyId string) ([]*SdkPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules", policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var policyRule []*SdkPolicyRule

	resp, err := rq.Do(ctx, req, &policyRule)
	if err != nil {
		return nil, resp, err
	}

	return policyRule, resp, nil
}

// Creates a policy rule.
func (m *PolicyResource) CreatePolicyRule(ctx context.Context, policyId string, body SdkPolicyRule) (*SdkPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules", policyId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policyRule *SdkPolicyRule

	resp, err := rq.Do(ctx, req, &policyRule)
	if err != nil {
		return nil, resp, err
	}

	return policyRule, resp, nil
}

// Removes a policy rule.
func (m *PolicyResource) DeletePolicyRule(ctx context.Context, policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Gets a policy rule.
func (m *PolicyResource) GetPolicyRule(ctx context.Context, policyId string, ruleId string) (*SdkPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var policyRule *SdkPolicyRule

	resp, err := rq.Do(ctx, req, &policyRule)
	if err != nil {
		return nil, resp, err
	}

	return policyRule, resp, nil
}

// Updates a policy rule.
func (m *PolicyResource) UpdatePolicyRule(ctx context.Context, policyId string, ruleId string, body SdkPolicyRule) (*SdkPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v", policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var policyRule *SdkPolicyRule

	resp, err := rq.Do(ctx, req, &policyRule)
	if err != nil {
		return nil, resp, err
	}

	return policyRule, resp, nil
}

// Activates a policy rule.
func (m *PolicyResource) ActivatePolicyRule(ctx context.Context, policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v/lifecycle/activate", policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Deactivates a policy rule.
func (m *PolicyResource) DeactivatePolicyRule(ctx context.Context, policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/policies/%v/rules/%v/lifecycle/deactivate", policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (a *Policy) MarshalJSON() ([]byte, error) {
	type Alias Policy
	type local struct {
		*Alias
	}
	result := local{Alias: (*Alias)(a)}
	if a.Priority != 0 {
		result.PriorityPtr = Int64Ptr(a.Priority)
	}
	return json.Marshal(&result)
}

func (a *Policy) UnmarshalJSON(data []byte) error {
	type Alias Policy

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
	return nil
}
