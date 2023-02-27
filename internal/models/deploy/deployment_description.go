package deploy

import "time"

type ConditionStatus string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

type DeploymentCondition struct {
	Type    string          `json:"type,omitempty"`
	Status  ConditionStatus `json:"status"`
	Message string          `json:"message"`
}

// DeploymentDescription 裁剪后的Deployment Spec
type DeploymentDescription struct {
	Id              uint      `json:"id"`              // deployment.ID
	Replicate       uint      `json:"replicate"`       // 副本数
	AvailablePods   uint      `json:"availablePods"`   // 可用副本数
	UnavailablePods uint      `json:"unavailablePods"` // 不可用副本数
	CreatedAt       time.Time `json:"createdAt"`       // 创建时间
	Conditions      []string  `json:"conditions"`      // 状态
}
