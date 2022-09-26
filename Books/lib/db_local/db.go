package db_local

import (
	"DigitalPayment/lib/DB"
	"DigitalPayment/lib/parameters"
	"fmt"
	"gorm.io/gorm"
)

var (
	DB_LOCAL *gorm.DB
	err      error
)

func InitDB() error {
	dns := DB.ConnectDatabase{
		Host:     parameters.ParamsService.DBHost,
		Port:     parameters.ParamsService.DBPort,
		User:     parameters.ParamsService.DBUser,
		Password: parameters.ParamsService.DBPassword,
		DBName:   parameters.ParamsService.DBName,
	}
	DB_LOCAL, err = dns.ConnectToDatabase()
	if err != nil {
		fmt.Printf("ERROR CONNECT TO DATABASE: %s\n", err.Error())
		return err
	}
	return nil
}

const (
	tablename_books = "public.books"
)

func FindBookById(db *gorm.DB, where map[string]interface{}) (*Book, error) {
	bookData := Book{}

	tx := db.Table(tablename_books).Where(where).Select("*").Limit(1).Scan(&bookData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("записи в БД %s не найдены", tablename_books)
	}
	return &bookData, nil
}
func FindBooks(db *gorm.DB, where map[string]interface{}) (*Books, error) {
	booksData := Books{}

	tx := db.Table(tablename_books).Select("*").Where(where).Scan(&booksData.Books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("записи в БД %s не найдены", tablename_books)
	}
	return &booksData, nil
}

//TODO проверить
func CreateBook(db *gorm.DB, book *Book) (*Book, error) {
	var err error

	//проверка, что такая книга уже есть
	tx, err := FindBooks(db, map[string]interface{}{
		"Name":      book.Name,
		"Author":    book.Author,
		"Publisher": book.Publisher,
	})

	if err == nil && tx.Books != nil {
		err = fmt.Errorf("книга '%s' уже существует", book.Name+" "+book.Genre)
		return nil, err
	}

	bookData := Book{
		Name:        book.Name,
		Genre:       book.Genre,
		Author:      book.Author,
		Publisher:   book.Publisher,
		AddedUser:   book.AddedUser,
		AddedTime:   book.AddedTime,
		Description: book.Description,
	}
	txx := db.Table(tablename_books).Create(&bookData)

	if bookData.Id == 0 || txx.Error != nil || txx.RowsAffected == 0 {
		err = fmt.Errorf("ошибка записи в БД Books %s", tablename_books)
		return nil, err
	}
	return &bookData, nil
}
func ChangeBookById(db *gorm.DB, where map[string]interface{}, update map[string]interface{}) (*Book, error) {
	bookData := Book{}

	tx, err := FindBookById(db, where)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		err = fmt.Errorf("книга %+v не найдена", where)
		return nil, err
	}

	txx := db.Model(bookData).Where(where).Updates(update)
	if txx.Error != nil || txx.RowsAffected == 0 {
		err = fmt.Errorf("ошибка обновления в БД Books %s", tablename_books)
		return nil, err
	}

	return &bookData, nil
}
func RemoveBookById(db *gorm.DB, book *Book) error {
	tx := db.Where("id = ?", book.Id).Delete(book)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("записи в БД %s не найдены", tablename_books)
	}
	return nil
}