package nacos

import (
	"fmt"
	"github/szpinc/nacosctl/pkg/editor"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	client := Client{
		Config: &NacosConfig{
			Addr:       "http://172.16.8.123:8848/nacos",
			ApiVersion: "v1",
		},
	}

	content, err := client.Get(ConfigGetOperation{
		NacosOperation: &DefaultNacosOperation,
		DataId:         "common.yaml",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(content)
}

func TestEdit(t *testing.T) {

	configFile := "/Users/ghostdog/GoProjects/nacos-cli/basic-data-webapi.yaml"

	e := editor.NewDefaultEditor([]string{})

	err := e.Launch(configFile)

	if err != nil {
		panic(err)
	}

	fmt.Println("Edited")

	contentBytes, err := os.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	client := Client{
		Config: &NacosConfig{
			Addr:       "http://172.16.8.123:8848/nacos",
			ApiVersion: "v1",
		},
	}

	err = client.Edit(ConfigEditOperation{
		NacosOperation: &DefaultNacosOperation,
		DataId:         "common.yaml",
		Content:        string(contentBytes),
	})

	if err != nil {
		panic(err)
	}

}

func TestDeleteConfig(t *testing.T) {
	client := Client{
		Config: &NacosConfig{
			Addr:       "http://172.16.8.123:8848/nacos",
			ApiVersion: "v1",
		},
	}

	err := client.DeleteConfig(ConfigDeleteOperation{
		NacosOperation: &NacosOperation{
			Group:     "DEFAULT_GROUP",
			Namespace: "demo",
		},
		DataId: "common.yaml",
	})

	assert.Nil(t, err)
}
