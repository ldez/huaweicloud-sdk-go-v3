package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEvsQuotaRequest struct {
}

func (o ShowEvsQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEvsQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowEvsQuotaRequest", string(data)}, " ")
}
