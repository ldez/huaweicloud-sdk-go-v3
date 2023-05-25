package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建实例请求体。
type CreateInstanceByEngineReq struct {

	// 实例名称。  由英文字符开头，只能由英文字母、数字、中划线、下划线组成，长度为4~64的字符。
	Name string `json:"name"`

	// 实例的描述信息。  长度不超过1024的字符串。  > \\与\"在json报文中属于特殊字符，如果参数值中需要显示\\或者\"字符，请在字符前增加转义字符\\，比如\\\\或者\\\"。
	Description *string `json:"description,omitempty"`

	// 消息引擎。取值填写为：kafka。
	Engine CreateInstanceByEngineReqEngine `json:"engine"`

	// 消息引擎的版本。取值填写为：   - 1.1.0   - 2.3.0   - 2.7
	EngineVersion CreateInstanceByEngineReqEngineVersion `json:"engine_version"`

	// 代理个数。
	BrokerNum int32 `json:"broker_num"`

	// 消息存储空间，单位GB。   - Kafka实例规格为c6.2u4g.cluster时，存储空间取值范围300GB ~ 300000GB。   - Kafka实例规格为c6.4u8g.cluster时，存储空间取值范围300GB ~ 600000GB。   - Kafka实例规格为c6.8u16g.cluster时，存储空间取值范围300GB ~ 900000GB。   - Kafka实例规格为c6.12u24g.cluster时，存储空间取值范围300GB ~ 900000GB。   - Kafka实例规格为c6.16u32g.cluster时，存储空间取值范围300GB ~ 900000GB。
	StorageSpace int32 `json:"storage_space"`

	// 当ssl_enable为true时，该参数必选，ssl_enable为false时，该参数无效。  认证用户名，只能由英文字母、数字、中划线组成，长度为4~64的字符。
	AccessUser *string `json:"access_user,omitempty"`

	// 当ssl_enable为true时，该参数必选，ssl_enable为false时，该参数无效。  实例的认证密码。  复杂度要求： - 输入长度为8到32位的字符串。 - 必须包含如下四种字符中的两种组合：   - 小写字母   - 大写字母   - 数字   - 特殊字符包括（`~!@#$%^&*()-_=+\\|[{}]:'\",<.>/?）
	Password *string `json:"password,omitempty"`

	// 虚拟私有云ID。  获取方法如下：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。
	VpcId string `json:"vpc_id"`

	// 指定实例所属的安全组。  获取方法如下：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	// 子网信息。  获取方法如下：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。
	SubnetId string `json:"subnet_id"`

	// 创建节点到指定且有资源的可用区ID。该参数不能为空数组或者数组的值为空。 创建Kafka实例，支持节点部署在1个或3个及3个以上的可用区。在为节点指定可用区时，用逗号分隔开。
	AvailableZones []string `json:"available_zones"`

	// 产品ID。 产品ID可以从**查询产品规格列表**接口查询到。
	ProductId string `json:"product_id"`

	// 表示登录Kafka Manager的用户名。只能由英文字母、数字、中划线组成，长度为4~64的字符。
	KafkaManagerUser *string `json:"kafka_manager_user,omitempty"`

	// 表示登录Kafka Manager的密码。  复杂度要求：   - 输入长度为8到32位的字符串。   - 必须包含如下四种字符中的两种组合：       - 小写字母       - 大写字母       - 数字       - 特殊字符包括（`~!@#$%^&*()-_=+\\|[{}]:'\",<.>/?）
	KafkaManagerPassword *string `json:"kafka_manager_password,omitempty"`

	// 维护时间窗开始时间，格式为HH:mm。 - 维护时间窗开始和结束时间必须为指定的时间段。 - 开始时间必须为22:00、02:00、06:00、10:00、14:00和18:00。 - 该参数不能单独为空，若该值为空，则结束时间也为空。系统分配一个默认开始时间02:00。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// 维护时间窗结束时间，格式为HH:mm。 - 维护时间窗开始和结束时间必须为指定的时间段。 - 结束时间在开始时间基础上加四个小时，即当开始时间为22:00时，结束时间为02:00。 - 该参数不能单独为空，若该值为空，则开始时间也为空，系统分配一个默认结束时间06:00。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// 是否开启公网访问功能。默认不开启公网。 - true：开启 - false：不开启
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// 实例绑定的弹性IP地址的ID。  以英文逗号隔开多个弹性IP地址的ID。  如果开启了公网访问功能（即enable_publicip为true），该字段为必选。
	PublicipId *string `json:"publicip_id,omitempty"`

	// 是否打开SSL加密访问。  实例创建后将不支持动态开启和关闭。  - true：打开SSL加密访问。 - false：不打开SSL加密访问。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// 开启SASL后使用的安全协议，如果开启了SASL认证功能（即ssl_enable=true），该字段为必选。  若该字段值为空，默认开启SASL_SSL认证机制。  实例创建后将不支持动态开启和关闭。  - SASL_SSL: 采用SSL证书进行加密传输，支持账号密码认证，安全性更高。 - SASL_PLAINTEXT: 明文传输，支持账号密码认证，性能更好，仅支持SCRAM-SHA-512机制。
	KafkaSecurityProtocol *string `json:"kafka_security_protocol,omitempty"`

	// 开启SASL后使用的认证机制，如果开启了SASL认证功能（即ssl_enable=true），该字段为必选。  若该字段值为空，默认开启PLAIN认证机制。  选择其一进行SASL认证即可，支持同时开启两种认证机制。 取值如下： - PLAIN: 简单的用户名密码校验。 - SCRAM-SHA-512: 用户凭证校验，安全性比PLAIN机制更高。
	SaslEnabledMechanisms *[]CreateInstanceByEngineReqSaslEnabledMechanisms `json:"sasl_enabled_mechanisms,omitempty"`

	// 磁盘的容量到达容量阈值后，对于消息的处理策略。  取值如下： - produce_reject：表示拒绝消息写入。 - time_base：表示自动删除最老消息。
	RetentionPolicy *CreateInstanceByEngineReqRetentionPolicy `json:"retention_policy,omitempty"`

	// 是否开启消息转储功能。  默认不开启消息转储。
	ConnectorEnable *bool `json:"connector_enable,omitempty"`

	// 是否打开kafka自动创建topic功能。 - true：开启 - false：关闭  当您选择开启，表示生产或消费一个未创建的Topic时，会自动创建一个包含3个分区和3个副本的Topic。  默认是false关闭。
	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty"`

	// 存储IO规格。  取值范围：   - dms.physical.storage.high.v2：使用高IO的磁盘类型。   - dms.physical.storage.ultra.v2：使用超高IO的磁盘类型。  如何选择磁盘类型请参考磁盘类型及性能介绍。
	StorageSpecCode CreateInstanceByEngineReqStorageSpecCode `json:"storage_spec_code"`

	// 企业项目ID。若为企业项目帐号，该参数必填。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签列表。
	Tags *[]TagEntity `json:"tags,omitempty"`

	// CPU架构。当前只支持X86架构。  取值范围：   - X86
	ArchType *string `json:"arch_type,omitempty"`

	// VPC内网明文访问。
	VpcClientPlain *bool `json:"vpc_client_plain,omitempty"`

	BssParam *BssParam `json:"bss_param,omitempty"`
}

func (o CreateInstanceByEngineReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceByEngineReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceByEngineReq", string(data)}, " ")
}

type CreateInstanceByEngineReqEngine struct {
	value string
}

type CreateInstanceByEngineReqEngineEnum struct {
	KAFKA CreateInstanceByEngineReqEngine
}

func GetCreateInstanceByEngineReqEngineEnum() CreateInstanceByEngineReqEngineEnum {
	return CreateInstanceByEngineReqEngineEnum{
		KAFKA: CreateInstanceByEngineReqEngine{
			value: "kafka",
		},
	}
}

func (c CreateInstanceByEngineReqEngine) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqEngine) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqEngineVersion struct {
	value string
}

type CreateInstanceByEngineReqEngineVersionEnum struct {
	E_1_1_0 CreateInstanceByEngineReqEngineVersion
	E_2_3_0 CreateInstanceByEngineReqEngineVersion
	E_2_7   CreateInstanceByEngineReqEngineVersion
}

func GetCreateInstanceByEngineReqEngineVersionEnum() CreateInstanceByEngineReqEngineVersionEnum {
	return CreateInstanceByEngineReqEngineVersionEnum{
		E_1_1_0: CreateInstanceByEngineReqEngineVersion{
			value: "1.1.0",
		},
		E_2_3_0: CreateInstanceByEngineReqEngineVersion{
			value: "2.3.0",
		},
		E_2_7: CreateInstanceByEngineReqEngineVersion{
			value: "2.7",
		},
	}
}

func (c CreateInstanceByEngineReqEngineVersion) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqEngineVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqEngineVersion) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqSaslEnabledMechanisms struct {
	value string
}

type CreateInstanceByEngineReqSaslEnabledMechanismsEnum struct {
	PLAIN         CreateInstanceByEngineReqSaslEnabledMechanisms
	SCRAM_SHA_512 CreateInstanceByEngineReqSaslEnabledMechanisms
}

func GetCreateInstanceByEngineReqSaslEnabledMechanismsEnum() CreateInstanceByEngineReqSaslEnabledMechanismsEnum {
	return CreateInstanceByEngineReqSaslEnabledMechanismsEnum{
		PLAIN: CreateInstanceByEngineReqSaslEnabledMechanisms{
			value: "PLAIN",
		},
		SCRAM_SHA_512: CreateInstanceByEngineReqSaslEnabledMechanisms{
			value: "SCRAM-SHA-512",
		},
	}
}

func (c CreateInstanceByEngineReqSaslEnabledMechanisms) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqSaslEnabledMechanisms) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqSaslEnabledMechanisms) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqRetentionPolicy struct {
	value string
}

type CreateInstanceByEngineReqRetentionPolicyEnum struct {
	TIME_BASE      CreateInstanceByEngineReqRetentionPolicy
	PRODUCE_REJECT CreateInstanceByEngineReqRetentionPolicy
}

func GetCreateInstanceByEngineReqRetentionPolicyEnum() CreateInstanceByEngineReqRetentionPolicyEnum {
	return CreateInstanceByEngineReqRetentionPolicyEnum{
		TIME_BASE: CreateInstanceByEngineReqRetentionPolicy{
			value: "time_base",
		},
		PRODUCE_REJECT: CreateInstanceByEngineReqRetentionPolicy{
			value: "produce_reject",
		},
	}
}

func (c CreateInstanceByEngineReqRetentionPolicy) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqRetentionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqRetentionPolicy) UnmarshalJSON(b []byte) error {
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

type CreateInstanceByEngineReqStorageSpecCode struct {
	value string
}

type CreateInstanceByEngineReqStorageSpecCodeEnum struct {
	DMS_PHYSICAL_STORAGE_HIGH_V2  CreateInstanceByEngineReqStorageSpecCode
	DMS_PHYSICAL_STORAGE_ULTRA_V2 CreateInstanceByEngineReqStorageSpecCode
}

func GetCreateInstanceByEngineReqStorageSpecCodeEnum() CreateInstanceByEngineReqStorageSpecCodeEnum {
	return CreateInstanceByEngineReqStorageSpecCodeEnum{
		DMS_PHYSICAL_STORAGE_HIGH_V2: CreateInstanceByEngineReqStorageSpecCode{
			value: "dms.physical.storage.high.v2",
		},
		DMS_PHYSICAL_STORAGE_ULTRA_V2: CreateInstanceByEngineReqStorageSpecCode{
			value: "dms.physical.storage.ultra.v2",
		},
	}
}

func (c CreateInstanceByEngineReqStorageSpecCode) Value() string {
	return c.value
}

func (c CreateInstanceByEngineReqStorageSpecCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineReqStorageSpecCode) UnmarshalJSON(b []byte) error {
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
