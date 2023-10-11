package bot

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"math/rand"
	"time"
	"untitledPetProject/internal/db"
)

var (
	helloMessage    = "Hello"
	byeMessage      = "Вы успешно отказались от участия в RandomCoffee от AK"
	newEventMessage = "Событие успешно создано"
)

func (h *handler) StartMessage(update *tgbot.Update) {
	user := &db.User{
		Id:       update.Message.Chat.ID,
		TGtag:    update.Message.Chat.UserName,
		IsActive: true,
	}
	if h.db.Model(user).Where("id = ?", update.Message.Chat.ID).Updates(user).RowsAffected == 0 {
		h.db.Create(user)
	}
	_, err := h.bot.Send(tgbot.NewMessage(update.Message.Chat.ID, helloMessage))
	if err != nil {
		log.Error(err)
	}
}

func (h *handler) Disactivate(update *tgbot.Update) {
	user := &db.User{
		Id:       update.Message.Chat.ID,
		TGtag:    update.Message.Chat.UserName,
		IsActive: false,
	}
	h.db.Save(&user)
	_, err := h.bot.Send(tgbot.NewMessage(user.Id, byeMessage))
	if err != nil {
		log.Error(err)
	}
}

func isAdmin(userid int64, database *gorm.DB) bool {
	var admin db.Admin
	result := database.First(&admin, userid)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func (h *handler) CreateEvent(update *tgbot.Update) {
	h.db.Model(&db.RCEvent{}).Where("is_active = ?", true).Update("is_active", false)
	event := db.RCEvent{
		DateStarted: time.Now(),
		IsActive:    true,
	}
	h.db.Create(&event)

	var users []db.User
	h.db.Find(&users, "is_active = ?", true)
	if len(users)%2 == 1 {
		_, err := h.bot.Send(tgbot.NewMessage(update.Message.Chat.ID, "Количество активных участников нечетное"))
		if err != nil {
			log.Error(err)
		}
		return
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(users), func(i, j int) {
		users[i], users[j] = users[j], users[i]
	})

	for i := 0; i < len(users); i += 2 {
		pair := db.Pair{
			EventId: event.Id,
			User1Id: users[i].Id,
			User2Id: users[i+1].Id,
		}
		h.db.Create(&pair)
	}
	var pairs []db.Pair
	h.db.Find(&pairs, "event_id = ?", event.Id)
	// log.Info(pairs)
	// TODO message to all pairs
	_, err := h.bot.Send(tgbot.NewMessage(update.Message.Chat.ID, newEventMessage))
	if err != nil {
		log.Error(err)
	}
}
