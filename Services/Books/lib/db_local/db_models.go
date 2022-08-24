package db_local

import "time"

type Book struct {
	Id          int64     `gorm:"primary key `
	Name        string    `gorm:"column:Name"`
	Genre       string    `gorm:"column:Genre"`
	Author      int64     `gorm:"column:Author"`
	Publisher   int64     `gorm:"column:Publisher"`
	AddedUser   int64     `gorm:"column:AddedUser"`
	AddedTime   time.Time `gorm:"column:AddedTime"`
	Description string    `gorm:"column:Description"`
}
type Books struct {
	Books []Book
}
