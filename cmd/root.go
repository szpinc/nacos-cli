package cmd

import (
	"github/szpinc/nacosctl/pkg/nacos"
	"os"

	"github.com/spf13/cobra"
)

var namespace string
var group string

var nacosClient *nacos.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nacosctl",
	Short: "nacos cli tools",
	Long:  `nacos cli tools`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	// PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

	// },
	ValidArgs: []string{"get", "delete", "edit"},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "nacos namespace")
	rootCmd.PersistentFlags().StringVarP(&group, "group", "g", "DEFAULT_GROUP", "nacos group")

	_ = rootCmd.MarkFlagRequired("namespace")
	_ = rootCmd.MarkFlagRequired("group")

	nacosClient = nacos.NewDefaultClient()
}
