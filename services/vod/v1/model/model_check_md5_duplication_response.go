package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CheckMd5DuplicationResponse struct {
	// 是否重复

	IsDuplicated *int32 `json:"is_duplicated,omitempty"`
	// 重复的媒资ID

	AssetIds       *[]string `json:"asset_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CheckMd5DuplicationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckMd5DuplicationResponse struct{}"
	}

	return strings.Join([]string{"CheckMd5DuplicationResponse", string(data)}, " ")
}
