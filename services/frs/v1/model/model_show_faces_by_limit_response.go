package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowFacesByLimitResponse struct {
	// 人脸库ID，随机生成的包含八个字符的字符串。 调用失败时无此字段。

	FaceSetId *string `json:"face_set_id,omitempty"`
	// 人脸库名称。 调用失败时无此字段。

	FaceSetName *string `json:"face_set_name,omitempty"`
	// 人脸库当中的人脸结构，详见[FaceSetFace](zh-cn_topic_0106912070.xml)。 调用失败时无此字段。

	Faces          *[]FaceSetFace `json:"faces,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowFacesByLimitResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowFacesByLimitResponse struct{}"
	}

	return strings.Join([]string{"ShowFacesByLimitResponse", string(data)}, " ")
}
