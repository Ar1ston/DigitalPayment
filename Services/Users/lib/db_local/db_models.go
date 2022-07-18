package db_local

type User struct {
	Id       int64  `gorm:"primary key"`
	Login    string `gorm:"Login"`
	Password string
	Name     string
	Level    int64
}
type Users struct {
	Users []User
}
