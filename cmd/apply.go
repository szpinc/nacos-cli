/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github/szpinc/nacosctl/pkg/nacos"

	"github.com/spf13/cobra"
)

var (
	file   string
	dataId string
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "create or update resource",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nacosClient.ApplyConfig(nacos.ConfigApplyOperation{
			NacosOperation: &nacos.NacosOperation{
				Namespace: namespace,
				Group:     group,
			},
			DataId: dataId,
			File:   file,
			Type:   fileType,
		})
	},
}

func init() {

	applyCmd.Flags().StringVarP(&file, "file", "f", "", "config file(required)")
	applyCmd.Flags().StringVarP(&dataId, "id", "d", "", "data id")
	applyCmd.Flags().StringVarP(&fileType, "type", "t", "", "config file type. e.g: yaml")

	applyCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(applyCmd)
}
