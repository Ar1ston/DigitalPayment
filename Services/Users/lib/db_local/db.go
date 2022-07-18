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
		DBName:   "users",
	}
	DB_LOCAL, err = dns.ConnectToDatabase()
	if err != nil {
		fmt.Printf("ERROR CONNECT TO DATABASE: %s\n", err.Error())
	}
}

const (
	tablename_users = "public.users"
)

func FindUser(db *gorm.DB, where map[string]interface{}) (*User, error) {
	userData := User{}

	tx := db.Table(tablename_users).Where(where).Select("*").Limit(1).Scan(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("записи в БД %s не найдены", tablename_users)
	}
	return &userData, nil
}
func FindUsers(db *gorm.DB, where map[string]interface{}) (*Users, error) {
	usersData := Users{}

	tx := db.Table(tablename_users).Select("*").Where(where).Scan(&usersData.Users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("%s%s%s", "Записи в БД ", tablename_users, " не найдены")
	}
	return &usersData, nil
}
func CreateUser(db *gorm.DB, user *User) (*User, error) {
	var err error

	//проверка, что такой пользователь уже есть
	tx, err := FindUsers(db, map[string]interface{}{
		"login": user.Login,
	})

	if err == nil && tx != nil {
		err = fmt.Errorf("пользователь '%s' уже существует", user.Login)
		return nil, err
	}

	userData := User{
		Login:    user.Login,
		Password: user.Password,
		Name:     user.Name,
	}
	txx := db.Table(tablename_users).Create(&userData)

	if userData.Id == 0 || txx.Error != nil || txx.RowsAffected == 0 {
		err = fmt.Errorf("ошибка записи в БД Users %s", tablename_users)
		return nil, err
	}
	return &userData, nil
}
func ChangeUserById(db *gorm.DB, where map[string]interface{}, update map[string]interface{}) (*User, error) {
	userData := User{}

	tx, err := FindUser(db, where)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		err = fmt.Errorf("пользователь %+v не найден", where)
		return nil, err
	}

	txx := db.Table(tablename_users).Model(userData).Where(where).Updates(update)
	if txx.Error != nil || txx.RowsAffected == 0 {
		err = fmt.Errorf("ошибка обновления в БД Users %s", tablename_users)
		return nil, err
	}

	return &userData, nil
}
func RemoveUserById(db *gorm.DB, user *User) error {
	tx := db.Table(tablename_users).Where("id = ?", user.Id).Delete(user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("записи в БД %s не найдены", tablename_users)
	}
	return nil
}
