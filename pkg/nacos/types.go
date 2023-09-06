package nacos

type NacosConfig struct {
	Addr       string `json:"addr" yaml:"addr"`
	Username   string `json:"username" yaml:"username"`
	Password   string `json:"password" yaml:"password"`
	ApiVersion string `json:"apiVersion" yaml:"apiVersion"`
}

type NacosOperation struct {
	Namespace string // 命名空间
	Group     string // 组
}

// ConfigEditOperation 配置更新操作
type ConfigEditOperation struct {
	*NacosOperation
	Content string // 配置内容
	DataId  string // data-id
	Type    string // 文件类型
}

// ConfigEditOperation 配置查询操作
type ConfigGetOperation struct {
	*NacosOperation
	DataId string // data-id
}

var DefaultNacosOperation = NacosOperation{
	Namespace: "public",
	Group:     "DEFAULT_GROUP",
}

type NacosPageResult struct {
	PageItems []NacosPageItem `json:"pageItems"`
}

type NacosPageItem struct {
	Id     string `json:"id"`
	DataId string `json:"dataId"`
	Group  string `json:"group"`
	Type   string `json:"type"`   // 文件类型
	Tenant string `json:"tenant"` // 命名空间
}

// NacosConfigDetail nacos配置结构体
type NacosConfigDetail struct {
	ID               string `json:"id"`
	DataID           string `json:"dataId"`
	Group            string `json:"group"`
	Content          string `json:"content"`
	Md5              string `json:"md5"`
	EncryptedDataKey string `json:"encryptedDataKey"`
	Tenant           string `json:"tenant"`
	AppName          string `json:"appName"`
	Type             string `json:"type"`
	CreateTime       int64  `json:"createTime"`
	ModifyTime       int64  `json:"modifyTime"`
	CreateUser       string `json:"createUser"`
	CreateIP         string `json:"createIp"`
	Desc             string `json:"desc"`
	Use              string `json:"use"`
	Effect           string `json:"effect"`
	Schema           string `json:"schema"`
}
