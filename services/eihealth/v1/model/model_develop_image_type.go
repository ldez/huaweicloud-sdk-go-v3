package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DevelopImageType struct {
	value string
}

type DevelopImageTypeEnum struct {
	SYSTEM   DevelopImageType
	CUSTOMER DevelopImageType
}

func GetDevelopImageTypeEnum() DevelopImageTypeEnum {
	return DevelopImageTypeEnum{
		SYSTEM: DevelopImageType{
			value: "SYSTEM",
		},
		CUSTOMER: DevelopImageType{
			value: "CUSTOMER",
		},
	}
}

func (c DevelopImageType) Value() string {
	return c.value
}

func (c DevelopImageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DevelopImageType) UnmarshalJSON(b []byte) error {
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
