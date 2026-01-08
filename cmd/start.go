package cmd

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Telegram bot",
	Run: func(cmd *cobra.Command, args []string) {
		runBot()
	},
}

func runBot() {
	token := os.Getenv("TELE_TOKEN")
	if token == "" {
		log.Fatal("TELE_TOKEN is not set")
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send("👋 Hi! I'm KVN.")
	})

	log.Println("infra-bot started")
	bot.Start()
}

func init() {
	rootCmd.AddCommand(startCmd)
}
