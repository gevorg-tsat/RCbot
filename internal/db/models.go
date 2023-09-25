package db

type Product struct {
	Id    int64 `gorm:"primarykey"`
	Code  string
	Price int64
}
