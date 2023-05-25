package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 存储库设置资源自动备份开关请求体
type VaultSetResourceReq struct {

	// 需要设置的资源id。
	ResourceIds []string `json:"resource_ids"`

	// 设置存储库资源动作
	Action VaultSetResourceReqAction `json:"action"`
}

func (o VaultSetResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultSetResourceReq struct{}"
	}

	return strings.Join([]string{"VaultSetResourceReq", string(data)}, " ")
}

type VaultSetResourceReqAction struct {
	value string
}

type VaultSetResourceReqActionEnum struct {
	SUSPEND   VaultSetResourceReqAction
	UNSUSPEND VaultSetResourceReqAction
}

func GetVaultSetResourceReqActionEnum() VaultSetResourceReqActionEnum {
	return VaultSetResourceReqActionEnum{
		SUSPEND: VaultSetResourceReqAction{
			value: "suspend",
		},
		UNSUSPEND: VaultSetResourceReqAction{
			value: "unsuspend",
		},
	}
}

func (c VaultSetResourceReqAction) Value() string {
	return c.value
}

func (c VaultSetResourceReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VaultSetResourceReqAction) UnmarshalJSON(b []byte) error {
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
