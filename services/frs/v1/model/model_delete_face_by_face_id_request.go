package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteFaceByFaceIdRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`
	// 人脸ID。

	FaceId string `json:"face_id"`
}

func (o DeleteFaceByFaceIdRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteFaceByFaceIdRequest struct{}"
	}

	return strings.Join([]string{"DeleteFaceByFaceIdRequest", string(data)}, " ")
}
