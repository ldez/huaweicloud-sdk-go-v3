package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateExtractAudioTaskResponse struct {
	// 视频源媒资ID<br/>

	AssetId *string `json:"asset_id,omitempty"`
	// 提取的音频媒资ID<br/>

	AudioAssetId   *string `json:"audio_asset_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateExtractAudioTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateExtractAudioTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateExtractAudioTaskResponse", string(data)}, " ")
}
