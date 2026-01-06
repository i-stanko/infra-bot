package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print infra-bot version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("infra-bot version:", appVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
