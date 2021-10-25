package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateProtectedInstanceNameRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *UpdateProtectedInstanceNameRequestBody `json:"body,omitempty"`
}

func (o UpdateProtectedInstanceNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateProtectedInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateProtectedInstanceNameRequest", string(data)}, " ")
}
