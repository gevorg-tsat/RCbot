package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id       uint64 `gorm:"primarykey"`
	TGtag    string
	IsActive bool
}

type RCEvent struct {
	gorm.Model
	Id          uint64 `gorm:"primarykey"`
	DateStarted time.Time
	IsActive    bool
}

type Pair struct {
	gorm.Model
	EventId uint64
	Event   RCEvent
	User1Id uint64
	User1   User
	User2Id uint64
	User2   User
}
