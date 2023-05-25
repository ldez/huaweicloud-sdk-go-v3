package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListExternalVaultRequest struct {

	// 其他区域的项目ID
	ExternalProjectId string `json:"external_project_id"`

	// 每页显示条目数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移值
	Offset *int32 `json:"offset,omitempty"`

	// [保护类型。取值为backup，replication和hybrid。](tag:hws,hws_hk) [保护类型。取值为backup和replication。](tag:ocb) [保护类型。取值为backup。](tag:g42,hk-g42,sbc,dt,fcs_vm,ctc,tm,tlf,cmcc,hcso_dt)
	ProtectType *ListExternalVaultRequestProtectType `json:"protect_type,omitempty"`

	// 区域ID
	RegionId string `json:"region_id"`

	// 资源类型
	ObjcetType *string `json:"objcet_type,omitempty"`

	// [云类型。取值为public和hybrid。](tag:hws,hws_hk) [云类型。取值为public。](tag:g42,hk-g42,sbc,dt,fcs_vm,ctc,ocb,tm,tlf,cmcc,hcso_dt)
	CloudType *string `json:"cloud_type,omitempty"`

	// 存储库ID，指定存储ID时其他过滤条件不生效。
	VaultId *string `json:"vault_id,omitempty"`
}

func (o ListExternalVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExternalVaultRequest struct{}"
	}

	return strings.Join([]string{"ListExternalVaultRequest", string(data)}, " ")
}

type ListExternalVaultRequestProtectType struct {
	value string
}

type ListExternalVaultRequestProtectTypeEnum struct {
	BACKUPREPLICATIONHYBRID ListExternalVaultRequestProtectType
}

func GetListExternalVaultRequestProtectTypeEnum() ListExternalVaultRequestProtectTypeEnum {
	return ListExternalVaultRequestProtectTypeEnum{
		BACKUPREPLICATIONHYBRID: ListExternalVaultRequestProtectType{
			value: "backup;replication;hybrid",
		},
	}
}

func (c ListExternalVaultRequestProtectType) Value() string {
	return c.value
}

func (c ListExternalVaultRequestProtectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListExternalVaultRequestProtectType) UnmarshalJSON(b []byte) error {
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
