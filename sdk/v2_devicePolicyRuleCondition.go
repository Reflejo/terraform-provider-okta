package sdk

type DevicePolicyRuleCondition struct {
	Migrated   *bool                              `json:"migrated,omitempty"`
	Platform   *DevicePolicyRuleConditionPlatform `json:"platform,omitempty"`
	Rooted     *bool                              `json:"rooted,omitempty"`
	TrustLevel string                             `json:"trustLevel,omitempty"`
}

func NewDevicePolicyRuleCondition() *DevicePolicyRuleCondition {
	return &DevicePolicyRuleCondition{}
}

func (a *DevicePolicyRuleCondition) IsPolicyInstance() bool {
	return true
}
