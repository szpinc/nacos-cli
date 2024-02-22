package nacos

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

// ApplyConfig 新增 or 修改配置
// configFile:
func (c *Client) ApplyConfig(operation ConfigApplyOperation) error {

	file, err := os.Open(operation.File)

	if err != nil {
		return err
	}

	buf, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	dataType := operation.Type

	if dataType == "" {
		dataType = strings.ReplaceAll(path.Ext(operation.File), ".", "")
	}

	if operation.DataId == "" {
		operation.DataId = path.Base(operation.File)
	}

	if err = c.Edit(ConfigEditOperation{
		NacosOperation: operation.NacosOperation,
		Content:        string(buf),
		DataId:         operation.DataId,
		Type:           dataType,
	}); err != nil {
		return err
	}
	fmt.Println("OK!")
	return nil
}

// DeleteConfig 删除配置
func (c *Client) DeleteConfig(operation ConfigDeleteOperation) error {
	deleteUrl, err := getUrl(c.Config)

	if err != nil {
		return err
	}

	resp, err := http.PostForm(deleteUrl, url.Values{
		"dataId": []string{operation.DataId},
		"group":  []string{operation.Group},
	})

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response error,status code:%d", resp.StatusCode)
	}

	return nil
}
