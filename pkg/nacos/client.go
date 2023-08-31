package nacos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	baseUrl = "/cs/configs"
)

// Client Nacos客户端
type Client struct {
	Config *NacosConfig
}

// Get获取配置
func (c *Client) Get(operation ConfigGetOperation) (string, error) {

	configUrl, err := getUrl(c.Config)

	if err != nil {
		return "", err
	}

	requestUrl := fmt.Sprintf(configUrl+"?dataId=%s&group=%s&tenant=%s", operation.DataId, operation.Group, operation.Namespace)

	resp, err := http.Get(requestUrl)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		body = []byte{}
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response error,status code:%d\n%s", resp.StatusCode, body)
	}

	return string(body), nil
}

// AllConfig 获取所有配置
func (c *Client) AllConfig(operation ConfigGetOperation) ([]NacosPageItem, error) {

	println("Namespace: ", operation.Namespace)

	configUrl, err := getUrl(c.Config)

	if err != nil {
		return nil, err
	}

	requestUrl := fmt.Sprintf(configUrl+"?dataId=&group=%s&tenant=%s&pageNo=1&pageSize=999&search=accurate", operation.Group, operation.Namespace)

	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error,status code:%d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		body = []byte{}
	}

	result := NacosPageResult{}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.PageItems, nil
}

// Edit 更新配置
func (c *Client) Edit(operation ConfigEditOperation) error {

	configUrl, err := getUrl(c.Config)

	if err != nil {
		return err
	}

	resp, err := http.PostForm(configUrl, url.Values{
		"dataId":  []string{operation.DataId},
		"group":   []string{operation.Group},
		"content": []string{operation.Content},
		"tenant":  []string{operation.Namespace},
	})

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response error,status code:%d", resp.StatusCode)
	}

	return nil
}

func getUrl(config *NacosConfig) (string, error) {
	return url.JoinPath(config.Addr, config.ApiVersion, baseUrl)
}

func NewDefaultClient() *Client {

	addr := os.Getenv("NACOS_ADDR")
	apiVersion := os.Getenv("NACOS_API_VERSION")

	if addr == "" {
		addr = "http://127.0.0.1:8848/nacos"
	}

	if apiVersion == "" {
		apiVersion = "v1"
	}

	return &Client{
		Config: &NacosConfig{
			Addr:       addr,
			ApiVersion: apiVersion,
		},
	}
}
