package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchAddTagsRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *BatchAddTagsRequestBody `json:"body,omitempty"`
}

func (o BatchAddTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchAddTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddTagsRequest", string(data)}, " ")
}
