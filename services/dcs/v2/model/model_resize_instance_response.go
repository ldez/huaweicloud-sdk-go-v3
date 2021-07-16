package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ResizeInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResizeInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResizeInstanceResponse struct{}"
	}

	return strings.Join([]string{"ResizeInstanceResponse", string(data)}, " ")
}
