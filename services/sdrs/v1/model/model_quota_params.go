package model

import (
	"encoding/json"

	"strings"
)

// 配额数据结构
type QuotaParams struct {
	// 租户资源配额信息列表。

	Resources []QuotaResourceParams `json:"resources"`
}

func (o QuotaParams) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QuotaParams struct{}"
	}

	return strings.Join([]string{"QuotaParams", string(data)}, " ")
}
