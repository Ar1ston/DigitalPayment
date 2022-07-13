package db_local

import (
	"DigitalPayment/lib/DB"
	"fmt"
	"gorm.io/gorm"
)

var (
	DB_LOCAL *gorm.DB
	err      error
)

func init() {
	dns := DB.ConnectDatabase{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "authors",
	}
	DB_LOCAL, err = dns.ConnectToDatabase()
	if err != nil {
		fmt.Printf("ERROR CONNECT TO DATABASE: %s\n", err.Error())
	}
}

const (
	tablename_authors = "public.authors"
)

type Author struct {
	Id          int64 `gorm:"primary key"`
	FirstName   string
	LastName    string
	Description string
}

func FindAuthorById(db *gorm.DB, where map[string]interface{}) (*Author, error) {

	authorsData := Author{}

	tx := db.Table(tablename_authors).Where(where).Select("*").Where(where).Limit(1).Scan(&authorsData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("%s%s%s", "Записи в БД ", tablename_authors, " не найдены")
	}
	return &authorsData, nil
}
