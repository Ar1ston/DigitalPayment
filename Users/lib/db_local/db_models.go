package db_local

type User struct {
	Id       int64  `gorm:"primary key"`
	Login    string `gorm:"column:Login"`
	Password string `gorm:"column:Password"`
	Name     string `gorm:"column:Name"`
	Level    int64  `gorm:"column:Level"`
}
type Users struct {
	Users []User
}
