package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddFacesByBase64Request struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`

	Body *AddFacesBase64Req `json:"body,omitempty"`
}

func (o AddFacesByBase64Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddFacesByBase64Request struct{}"
	}

	return strings.Join([]string{"AddFacesByBase64Request", string(data)}, " ")
}
