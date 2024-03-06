package common

// 真假
const (
	True  = 1 //是
	False = 0 //否
)

// 状态
const (
	Init   = 0 //待沟通
	PutOn  = 1 //上架
	PutOff = 2 //下架
)

// 图片类型
const (
	Logo               = 200 //logo
	BusinessLicense    = 201 //营业执照
	Wechat             = 202 //微信
	ServiceCase        = 203 //服务案例
	ExclusiveResources = 204 //独家资源
	CompanyPicture     = 205 //公司图片
)

// CompanyBasicsInfo 公司基本信息
const CompanyBasicsInfo = "company_basic_info"

// ServiceCapability 服务能力
const ServiceCapability = "service_capability"

// ServiceCases 服务案例
const ServiceCases = "service_case"

// ExclusiveResource 独家资源
const ExclusiveResource = "exclusive_resources"

// CompanyLabel 公司标签
const CompanyLabel = "company_label"
