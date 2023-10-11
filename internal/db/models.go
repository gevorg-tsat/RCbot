package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id       int64 `gorm:"primarykey"`
	TGtag    string
	IsActive bool
}

type RCEvent struct {
	gorm.Model
	Id          int64 `gorm:"primarykey"`
	DateStarted time.Time
	IsActive    bool
	Pairs       []Pair `gorm:"foreignKey:EventId"`
}

type Pair struct {
	gorm.Model
	EventId int64
	Event   RCEvent
	User1Id int64
	User1   User
	User2Id int64
	User2   User
}
type Admin struct {
	gorm.Model
	Id int64 `gorm:"primarykey"`
}
