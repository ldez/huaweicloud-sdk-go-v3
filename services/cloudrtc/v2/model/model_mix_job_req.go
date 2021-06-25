package model

import (
	"encoding/json"

	"strings"
)

// 合流任务参数，转推和录制至少选一个
type MixJobReq struct {
	MixParam *MixParam `json:"mix_param"`

	PublishParam *PublishParam `json:"publish_param,omitempty"`

	RecordParam *RecordParam `json:"record_param,omitempty"`
}

func (o MixJobReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MixJobReq struct{}"
	}

	return strings.Join([]string{"MixJobReq", string(data)}, " ")
}
