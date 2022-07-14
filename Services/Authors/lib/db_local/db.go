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

func FindAuthorById(db *gorm.DB, where map[string]interface{}) (*Author, error) {

	authorData := Author{}

	tx := db.Table(tablename_authors).Where(where).Select("*").Where(where).Limit(1).Scan(&authorData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("%s%s%s", "Записи в БД ", tablename_authors, " не найдены")
	}
	return &authorData, nil
}
func FindAuthors(db *gorm.DB, where map[string]interface{}) (*Authors, error) {
	authorsData := Authors{}

	tx := db.Table(tablename_authors).Where(where).Select("*").Where(where).Scan(&authorsData.Authors)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("%s%s%s", "Записи в БД ", tablename_authors, " не найдены")
	}
	return &authorsData, nil
}
func CreateAuthor(db *gorm.DB, first_name string, last_name string, description string) (int, error) {
	var err error

	//проверка, что такое автор уже есть
	tx, err := FindAuthors(db, map[string]interface{}{
		"FirstName": first_name,
		"LastName":  last_name,
	})

	if err == nil && tx != nil {
		err = fmt.Errorf("автор '%s' уже существует", first_name+" "+last_name)
		return 0, err
	}

	authorData := Author{
		FirstName:   first_name,
		LastName:    last_name,
		Description: description,
	}
	txx := db.Create(&authorData)

	if authorData.Id == 0 || txx.Error != nil || txx.RowsAffected == 0 {
		err = fmt.Errorf("ошибка записи в БД Authors %s", tablename_authors)
		return 0, err
	}
	return int(authorData.Id), nil
}
