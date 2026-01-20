package bot

import (
	"context"
	"log"
	"time"

	telebot "gopkg.in/telebot.v4"
	"github.com/i-stanko/infra-bot/internal/config"
)

func Run(ctx context.Context, cfg *config.Config, dryRun bool) error {
	if dryRun {
		log.Println("[Dry-Run] Bot would start here")
		return nil
	}

	b, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.TelegramToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return err
	}

	b.Handle(telebot.OnText, func(c telebot.Context) error {
		return c.Send("👋 Hi! I'm KVN.")
	})

	log.Println("infra-bot started")
	go b.Start()

	<-ctx.Done()
	log.Println("infra-bot stopped")
	b.Stop()
	return nil
}
