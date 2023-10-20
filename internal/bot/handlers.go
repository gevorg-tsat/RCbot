package bot

import (
	"RCbot/internal/db"
	"fmt"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	helloMessage    = "Привет! Добро пожаловать в RandomCoffee от АК. Совсем скоро мы напишем тебе о том, с кем тебе предстоит встретиться!\nЕсли вдруг захочешь прекратить участие -- введи команду /stop"
	byeMessage      = "Вы успешно отказались от участия в RandomCoffee от AK"
	newEventMessage = "Событие успешно создано"
	newPair         = "Привет! На предстоящей неделе твой партнер по RandomCoffee это %v. Договоритесь про место и время!"
)

func (h *handler) StartMessage(update *tgbot.Update) {
	user := &db.User{
		Id:       update.Message.Chat.ID,
		TGtag:    update.Message.Chat.UserName,
		IsActive: true,
	}
	if h.db.Model(user).Where("id = ?", update.Message.Chat.ID).Updates(user).RowsAffected == 0 {
		h.db.Create(&user)
	}
	_, err := h.bot.Send(tgbot.NewMessage(update.Message.Chat.ID, helloMessage))
	if err != nil {
		log.Error(err)
	}
	var adm db.Admin
	var cnt int64
	if h.db.Model(adm).Count(&cnt); cnt == 0 {
		if id, err := strconv.ParseInt(os.Getenv("FIRST_ADMIN_TG_ID"), 10, 64); err == nil && id == update.Message.Chat.ID {
			adm.Id = id
			h.db.Create(&adm)
		}
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
	result := database.Find(&admin, userid)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func (h *handler) CreateEvent(update *tgbot.Update) {
	if !isAdmin(update.Message.Chat.ID, h.db) {
		return
	}
	h.db.Model(&db.RCEvent{}).Where("is_active = ?", true).Update("is_active", false)
	event := db.RCEvent{
		DateStarted: time.Now(),
		IsActive:    true,
	}
	var users []db.User
	h.db.Find(&users, "is_active = ?", true)
	if len(users)%2 == 1 {
		_, err := h.bot.Send(tgbot.NewMessage(update.Message.Chat.ID, "Количество активных участников нечетное"))
		if err != nil {
			log.Error(err)
		}
		return
	}
	h.db.Create(&event)
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
	for i := range pairs {
		_, err := h.bot.Send(tgbot.NewMessage(pairs[i].User1Id, fmt.Sprintf(newPair, pairs[i].User2.TGtag)))
		if err != nil {
			h.bot.Send(tgbot.NewMessage(update.Message.Chat.ID, err.Error()))
		}
		_, err = h.bot.Send(tgbot.NewMessage(pairs[i].User2Id, fmt.Sprintf(newPair, pairs[i].User1.TGtag)))
		if err != nil {
			h.bot.Send(tgbot.NewMessage(update.Message.Chat.ID, err.Error()))
		}
	}
	_, err := h.bot.Send(tgbot.NewMessage(update.Message.Chat.ID, newEventMessage))
	if err != nil {
		log.Error(err)
	}
}
