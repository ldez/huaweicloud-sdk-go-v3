package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJobStatusRequest struct {
	// job ID

	JobId string `json:"job_id"`
}

func (o ShowJobStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowJobStatusRequest", string(data)}, " ")
}
