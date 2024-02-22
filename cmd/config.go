/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"github/szpinc/nacosctl/pkg/editor"
	"github/szpinc/nacosctl/pkg/nacos"
	"github/szpinc/nacosctl/pkg/util"
	"os"
	"path/filepath"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

var (
	getAllConfig bool   // 获取所有配置
	fileType     string // 配置类型
)

var getConfig = &cobra.Command{
	Use:   "config",
	Short: "nacos config",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		if getAllConfig {
			dataIds, err := nacosClient.AllConfig(nacos.ConfigGetOperation{
				NacosOperation: &nacos.NacosOperation{
					Namespace: namespace,
				},
			})

			if err != nil {
				return err
			}

			printTable(dataIds)
			return nil
		}

		if len(args) == 0 {
			return errors.New("data id required")
		}

		dataId := args[0]

		configData, err := nacosClient.Get(nacos.ConfigGetOperation{
			NacosOperation: &nacos.NacosOperation{
				Namespace: namespace,
				Group:     group,
			},
			DataId: dataId,
		})

		if err != nil {
			return err
		}

		fmt.Println(configData.Content)
		return nil
	},
}

var editConfig = &cobra.Command{
	Use:   "config",
	Short: "nacos config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		var dataId = args[0]

		configData, err := nacosClient.Get(nacos.ConfigGetOperation{
			NacosOperation: &nacos.NacosOperation{
				Namespace: namespace,
				Group:     group,
			},
			DataId: dataId,
		})

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		e := editor.NewDefaultEditor([]string{})

		buf := &bytes.Buffer{}
		buf.Write([]byte(configData.Content))

		edited, file, err := e.LaunchTempFile(fmt.Sprintf("%s-edit-", filepath.Base(os.Args[0])), configData.Type, buf)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		editedMd5 := util.Md5BytesToString(edited)

		if configData.Md5 == editedMd5 {
			fmt.Println("Not Changed")
			return
		}

		defer func(f string) {
			if e := os.Remove(f); e != nil {
				fmt.Println("delete temp file error:", e)
			}
		}(file)

		if fileType == "" {
			fileType = configData.Type
		}

		err = nacosClient.Edit(nacos.ConfigEditOperation{
			NacosOperation: &nacos.NacosOperation{
				Namespace: namespace,
				Group:     group,
			},
			DataId:  dataId,
			Content: string(edited),
			Type:    fileType,
		})

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Edited")
	},
}

var deleteConfig = &cobra.Command{
	Use:   "config",
	Short: "nacos config",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		nacosClient.DeleteConfig(nacos.ConfigDeleteOperation{
			NacosOperation: &nacos.NacosOperation{
				Namespace: namespace,
				Group:     group,
			},
			DataId: dataId,
		})
		return nil
	},
}

func init() {

	editConfig.Flags().StringVarP(&fileType, "type", "t", "", "file type")

	getConfig.Flags().BoolVarP(&getAllConfig, "all", "A", false, "If present, list the requested object(s) across all config name")

	editCmd.AddCommand(editConfig)
	getCmd.AddCommand(getConfig)
	deleteCmd.AddCommand(deleteConfig)
}

func printTable(items []nacos.NacosPageItem) {
	table := uitable.New()
	table.MaxColWidth = 50

	table.AddRow("ID", "GROUP", "NAMESPACE")

	for _, item := range items {
		if item.Tenant == "" {
			item.Tenant = "public"
		}
		table.AddRow(item.DataId, item.Group, item.Tenant)
	}

	fmt.Println(table)
}
