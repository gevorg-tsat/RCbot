package bot

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
	"os"
)

func Init() (*tgbot.BotAPI, error) {
	TOKEN := os.Getenv("BOT_TOKEN")
	log.Info("TOKEN: ", TOKEN)
	bot, err := tgbot.NewBotAPI(TOKEN)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	return bot, nil
}
