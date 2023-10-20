package bot

import (
	database "RCbot/internal/db"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
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

type handler struct {
    bot *tgbot.BotAPI
    db  *gorm.DB
}

func Run() error {
    db, err := database.Connect()
    if err != nil {
        return err
    }
    log.Info("Connected to Database")
    bot, err := createBot()
    log.Info("Bot connected")
    if err != nil {
        return err
    }
    updateConfig := tgbot.NewUpdate(0)
    updateConfig.Timeout = 60
    commandHandler := handler{bot: bot, db: db}
    for update := range bot.GetUpdatesChan(updateConfig) {
        command := update.Message.Command()
        if command == "start" {
            go commandHandler.StartMessage(&update)
        } else if command == "stop" {
            go commandHandler.Disactivate(&update)
        } else if command == "start_event" && isAdmin(update.Message.Chat.ID, db) {
            go commandHandler.CreateEvent(&update)
        }
    }
    return nil
}
