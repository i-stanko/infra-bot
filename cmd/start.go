package cmd

import (
	"context"
	"log"

	"github.com/i-stanko/infra-bot/internal/bot"
	"github.com/i-stanko/infra-bot/internal/config"
	"github.com/spf13/cobra"
)

var dryRun bool

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Telegram bot",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg *config.Config
		var err error

		if dryRun {
			cfg = &config.Config{}
		} else {
			cfg, err = config.LoadConfig()
			if err != nil {
				log.Fatal(err)
			}
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if err := bot.Run(ctx, cfg, dryRun); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	startCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Run bot in dry-run mode (no Telegram connection)")
	rootCmd.AddCommand(startCmd)
}
