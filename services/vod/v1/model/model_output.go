package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 转码输出数组。HLS或DASH场景，此数组的成员个数为n+1，n为转码输出路数。MP4场景，此数组的成员个数为n，n为转码输出路数
type Output struct {
	// 协议类型。 取值hls、 dash、mp4、mp3、aac

	PlayType OutputPlayType `json:"play_type"`
	// 访问URL

	Url string `json:"url"`
	// 标记流是否已被加密，取值[0,1] 0表示未加密，1表示已被加密。

	Encrypted *OutputEncrypted `json:"encrypted,omitempty"`
	// 清晰度

	Quality *OutputQuality `json:"quality,omitempty"`

	MetaData *MetaData `json:"meta_data"`
}

func (o Output) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Output struct{}"
	}

	return strings.Join([]string{"Output", string(data)}, " ")
}

type OutputPlayType struct {
	value string
}

type OutputPlayTypeEnum struct {
	HLS  OutputPlayType
	DASH OutputPlayType
	MP4  OutputPlayType
	MP3  OutputPlayType
	AAC  OutputPlayType
}

func GetOutputPlayTypeEnum() OutputPlayTypeEnum {
	return OutputPlayTypeEnum{
		HLS: OutputPlayType{
			value: "HLS",
		},
		DASH: OutputPlayType{
			value: "DASH",
		},
		MP4: OutputPlayType{
			value: "MP4",
		},
		MP3: OutputPlayType{
			value: "MP3",
		},
		AAC: OutputPlayType{
			value: "AAC",
		},
	}
}

func (c OutputPlayType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *OutputPlayType) UnmarshalJSON(b []byte) error {
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

type OutputEncrypted struct {
	value int32
}

type OutputEncryptedEnum struct {
	E_0 OutputEncrypted
	E_1 OutputEncrypted
}

func GetOutputEncryptedEnum() OutputEncryptedEnum {
	return OutputEncryptedEnum{
		E_0: OutputEncrypted{
			value: 0,
		}, E_1: OutputEncrypted{
			value: 1,
		},
	}
}

func (c OutputEncrypted) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *OutputEncrypted) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type OutputQuality struct {
	value string
}

type OutputQualityEnum struct {
	FLUENT  OutputQuality
	SD      OutputQuality
	HD      OutputQuality
	FULL_HD OutputQuality
}

func GetOutputQualityEnum() OutputQualityEnum {
	return OutputQualityEnum{
		FLUENT: OutputQuality{
			value: "FLUENT",
		},
		SD: OutputQuality{
			value: "SD",
		},
		HD: OutputQuality{
			value: "HD",
		},
		FULL_HD: OutputQuality{
			value: "FULL_HD",
		},
	}
}

func (c OutputQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *OutputQuality) UnmarshalJSON(b []byte) error {
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
