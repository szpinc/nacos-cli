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
}
