package bot

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
	"os"
)

func createBot() (*tgbot.BotAPI, error) {
	TOKEN := os.Getenv("BOT_TOKEN")
	bot, err := tgbot.NewBotAPI(TOKEN)
	if err != nil {
		return nil, err
	}
	//bot.Debug = true
	return bot, nil
}

func Run() error {
	bot, err := createBot()
	log.Info("Bot connected")
	if err != nil {
		return err
	}
	updateConfig := tgbot.NewUpdate(0)
	updateConfig.Timeout = 60
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message.Command() == "start" {
			go func() {
				_, err = bot.Send(tgbot.NewMessage(update.Message.Chat.ID, "Hello"))
				if err != nil {
					log.Error(err)
				}
			}()
		}
	}
	return nil
}
