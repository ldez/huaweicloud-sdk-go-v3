package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateTempResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTempResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTempResponse struct{}"
	}

	return strings.Join([]string{"UpdateTempResponse", string(data)}, " ")
}
