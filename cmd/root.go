package cmd

import (
	"github/szpinc/nacosctl/pkg/nacos"
	"os"

	"github.com/spf13/cobra"
)

var namespace string
var group string
var dataId string

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
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "PUBLIC", "nacos namespace")
	rootCmd.PersistentFlags().StringVarP(&group, "group", "g", "DEFAULT_GROUP", "nacos group")
	rootCmd.PersistentFlags().StringVarP(&dataId, "dataId", "d", "", "nacos dataId")

	_ = rootCmd.MarkFlagRequired("namespace")
	_ = rootCmd.MarkFlagRequired("group")

	nacosClient = nacos.NewDefaultClient()
}
