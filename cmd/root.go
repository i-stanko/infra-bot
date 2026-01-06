package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var appVersion = "dev"

var rootCmd = &cobra.Command{
	Use:   "infra-bot",
	Short: "Personal infrastructure automation and monitoring CLI",
	Long: `infra-bot is a CLI tool for running and managing a personal
infrastructure automation bot, including integrations with Telegram
and monitoring systems.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
