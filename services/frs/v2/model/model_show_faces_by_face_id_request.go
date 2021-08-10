package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowFacesByFaceIdRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`
	// 人脸ID。

	FaceId string `json:"face_id"`
}

func (o ShowFacesByFaceIdRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowFacesByFaceIdRequest struct{}"
	}

	return strings.Join([]string{"ShowFacesByFaceIdRequest", string(data)}, " ")
}
