/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"github/szpinc/nacosctl/pkg/editor"
	"github/szpinc/nacosctl/pkg/nacos"
	"github/szpinc/nacosctl/pkg/util"
	"os"
	"path/filepath"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

var getAllConfig bool

var getConfig = &cobra.Command{
	Use:   "config",
	Short: "nacos config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if dataId == "" && len(args) > 0 {
			dataId = args[0]
		}

		cfg := nacos.ConfigGetOperation{
			NacosOperation: &nacos.NacosOperation{
				Namespace: namespace,
				Group:     group,
			},
			DataId: dataId,
		}

		if dataId == "" {
			if getAllConfig {
				cfg.Group = ""
			}
			dataIds, err := nacosClient.AllConfig(cfg)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			printTable(dataIds)

			return
		}

		configData, err := nacosClient.Get(cfg)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(configData)
	},
}

var editConfig = &cobra.Command{
	Use:   "config",
	Short: "nacos config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if dataId == "" {
			dataId = args[0]
		}

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

		md5 := util.Md5ToString(configData)

		e := editor.NewDefaultEditor([]string{})

		buf := &bytes.Buffer{}
		buf.Write([]byte(configData))

		edited, file, err := e.LaunchTempFile(fmt.Sprintf("%s-edit-", filepath.Base(os.Args[0])), "yaml", buf)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		editedMd5 := util.Md5BytesToString(edited)

		if md5 == editedMd5 {
			fmt.Println("Not Changed")
			return
		}

		defer func(f string) {
			if e := os.Remove(f); e != nil {
				fmt.Println("delete temp file error:", e)
			}
		}(file)

		err = nacosClient.Edit(nacos.ConfigEditOperation{
			NacosOperation: &nacos.NacosOperation{
				Namespace: namespace,
				Group:     group,
			},
			DataId:  dataId,
			Content: string(edited),
		})

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Edited")
	},
}

func init() {
	editCmd.AddCommand(editConfig)
	getCmd.AddCommand(getConfig)

	getConfig.Flags().BoolVarP(&getAllConfig, "all", "A", false, "If present, list the requested object(s) across all config name")
}

func printTable(items []nacos.NacosPageItem) {
	table := uitable.New()
	table.MaxColWidth = 50

	table.AddRow("ID", "GROUP", "NAMESPACE")

	for _, item := range items {
		table.AddRow(item.DataId, item.Group, namespace)
	}

	fmt.Println(table)
}
