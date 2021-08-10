package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DetectLiveByFileResponse struct {
	VideoResult *LiveDetectRespVideoresult `json:"video-result,omitempty"`
	// 警告信息列表，WarningList结构见[WarningList](zh-cn_topic_0146322261.xml)。 调用失败时无此字段

	WarningList    *[]WarningList `json:"warning-list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o DetectLiveByFileResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectLiveByFileResponse struct{}"
	}

	return strings.Join([]string{"DetectLiveByFileResponse", string(data)}, " ")
}
