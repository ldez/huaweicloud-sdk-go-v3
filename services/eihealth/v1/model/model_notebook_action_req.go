package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 启停notebook请求体
type NotebookActionReq struct {
	Action *ActionType `json:"action"`
}

func (o NotebookActionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotebookActionReq struct{}"
	}

	return strings.Join([]string{"NotebookActionReq", string(data)}, " ")
}
