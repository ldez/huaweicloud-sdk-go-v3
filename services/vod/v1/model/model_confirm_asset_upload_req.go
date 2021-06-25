package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConfirmAssetUploadReq struct {
	// 媒体ID<br/>

	AssetId string `json:"asset_id"`
	// 上传状态<br/>

	Status ConfirmAssetUploadReqStatus `json:"status"`
}

func (o ConfirmAssetUploadReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConfirmAssetUploadReq struct{}"
	}

	return strings.Join([]string{"ConfirmAssetUploadReq", string(data)}, " ")
}

type ConfirmAssetUploadReqStatus struct {
	value string
}

type ConfirmAssetUploadReqStatusEnum struct {
	CREATED   ConfirmAssetUploadReqStatus
	FAILED    ConfirmAssetUploadReqStatus
	CANCELLED ConfirmAssetUploadReqStatus
}

func GetConfirmAssetUploadReqStatusEnum() ConfirmAssetUploadReqStatusEnum {
	return ConfirmAssetUploadReqStatusEnum{
		CREATED: ConfirmAssetUploadReqStatus{
			value: "CREATED",
		},
		FAILED: ConfirmAssetUploadReqStatus{
			value: "FAILED",
		},
		CANCELLED: ConfirmAssetUploadReqStatus{
			value: "CANCELLED",
		},
	}
}

func (c ConfirmAssetUploadReqStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ConfirmAssetUploadReqStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
