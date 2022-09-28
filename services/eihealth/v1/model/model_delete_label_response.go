package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteLabelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLabelResponse struct{}"
	}

	return strings.Join([]string{"DeleteLabelResponse", string(data)}, " ")
}
