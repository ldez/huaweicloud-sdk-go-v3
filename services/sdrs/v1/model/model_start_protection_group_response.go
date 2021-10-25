package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartProtectionGroupResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartProtectionGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartProtectionGroupResponse struct{}"
	}

	return strings.Join([]string{"StartProtectionGroupResponse", string(data)}, " ")
}
