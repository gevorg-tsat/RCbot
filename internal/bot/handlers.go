package bot

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
	"untitledPetProject/internal/db"
)

func (h *handler) StartMessage(update *tgbot.Update) {
	_, err := h.bot.Send(tgbot.NewMessage(update.Message.Chat.ID, "Hello"))
	if err != nil {
		log.Error(err)
	}
	user := &db.User{
		Id:       update.Message.Chat.ID,
		TGtag:    update.Message.Chat.UserName,
		IsActive: true,
	}
	if h.db.Model(&user).Where("id = ?", update.Message.Chat.ID).Updates(&user).RowsAffected == 0 {
		h.db.Create(&user)
	}
}
