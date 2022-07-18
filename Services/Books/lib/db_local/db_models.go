package db_local

import "time"

type Book struct {
	Id          int64 `gorm:"primary key"`
	Name        string
	Genre       string
	Author      int64
	Publisher   int64
	AddedUser   int64
	AddedTime   time.Time
	Description string
}
type Books struct {
	Books []Book
}
