package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateTemplateGroupResponse struct {
	// 模板组ID<br/>

	GroupId        *string `json:"group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTemplateGroupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTemplateGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateTemplateGroupResponse", string(data)}, " ")
}
