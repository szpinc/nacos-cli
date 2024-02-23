/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit a resource on the nacos server",
	Example: "nacosctl edit config common.yml",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	ValidArgs: []string{"config"},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
