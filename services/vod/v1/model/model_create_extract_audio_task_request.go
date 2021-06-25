package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateExtractAudioTaskRequest struct {
	Body *ExtractAudioTaskReq `json:"body,omitempty"`
}

func (o CreateExtractAudioTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateExtractAudioTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateExtractAudioTaskRequest", string(data)}, " ")
}
