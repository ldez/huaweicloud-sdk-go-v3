package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type SearchFaceByUrlResponse struct {
	// 查找的人脸集合，详见[SearchFace](zh-cn_topic_0106912071.xml)。 调用失败时无此字段。

	Faces          *[]SearchFace `json:"faces,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o SearchFaceByUrlResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SearchFaceByUrlResponse struct{}"
	}

	return strings.Join([]string{"SearchFaceByUrlResponse", string(data)}, " ")
}
