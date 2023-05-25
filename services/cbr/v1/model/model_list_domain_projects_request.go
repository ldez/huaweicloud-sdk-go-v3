package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDomainProjectsRequest struct {

	// 租户名称
	DomainName string `json:"domain_name"`
}

func (o ListDomainProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainProjectsRequest struct{}"
	}

	return strings.Join([]string{"ListDomainProjectsRequest", string(data)}, " ")
}
