package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSecretStageRequest struct {
	SecretId string `json:"secret_id"`

	StageName string `json:"stage_name"`
}

func (o ListSecretStageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSecretStageRequest struct{}"
	}

	return strings.Join([]string{"ListSecretStageRequest", string(data)}, " ")
}
