/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a resource from nacos server",
	Long: `
		Get a resource from nacos server
	`,
	Example: "nacosctl edit config common.yml",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	// ValidArgs: []string{"config"},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
