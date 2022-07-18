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
		DBName:   "publishers",
	}
	DB_LOCAL, err = dns.ConnectToDatabase()
	if err != nil {
		fmt.Printf("ERROR CONNECT TO DATABASE: %s\n", err.Error())
	}
}

const (
	tablename_publishers = "public.publishers"
)

func FindPublisherById(db *gorm.DB, where map[string]interface{}) (*Publisher, error) {
	publisherData := Publisher{}

	tx := db.Table(tablename_publishers).Where(where).Select("*").Limit(1).Scan(&publisherData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("записи в БД %s не найдены", tablename_publishers)
	}
	return &publisherData, nil
}
func FindPublishers(db *gorm.DB, where map[string]interface{}) (*Publishers, error) {
	publishersData := Publishers{}

	tx := db.Table(tablename_publishers).Select("*").Where(where).Scan(&publishersData.Publishers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("%s%s%s", "Записи в БД ", tablename_publishers, " не найдены")
	}
	return &publishersData, nil
}

//TODO проверить
func CreatePublisher(db *gorm.DB, publisher *Publisher) (*Publisher, error) {
	var err error

	//проверка, что такой издатель уже есть
	tx, err := FindPublishers(db, map[string]interface{}{
		"name": publisher.Name,
	})

	if err == nil && tx != nil {
		err = fmt.Errorf("издатель '%s' уже существует", publisher.Name)
		return nil, err
	}

	publisherData := Publisher{
		Name:        publisher.Name,
		Description: publisher.Description,
	}
	txx := db.Table(tablename_publishers).Create(&publisherData)

	if publisherData.Id == 0 || txx.Error != nil || txx.RowsAffected == 0 {
		err = fmt.Errorf("ошибка записи в БД Publishers %s", tablename_publishers)
		return nil, err
	}
	return &publisherData, nil
}
func ChangePublisherById(db *gorm.DB, where map[string]interface{}, update map[string]interface{}) (*Publisher, error) {
	publisherData := Publisher{}

	tx, err := FindPublisherById(db, where)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		err = fmt.Errorf("издатель %+v не найден", where)
		return nil, err
	}

	txx := db.Table(tablename_publishers).Model(publisherData).Where(where).Updates(update)
	if txx.Error != nil || txx.RowsAffected == 0 {
		err = fmt.Errorf("ошибка обновления в БД Publishers %s", tablename_publishers)
		return nil, err
	}

	return &publisherData, nil
}
func RemovePublisherById(db *gorm.DB, publisher *Publisher) error {
	tx := db.Where("id = ?", publisher.Id).Delete(publisher)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("записи в БД %s не найдены", tablename_publishers)
	}
	return nil
}
