package db_local

type Publisher struct {
	Id          int64  `gorm:"primary key"`
	Name        string `gorm:"column:Name"`
	Description string `gorm:"column:Description"`
}
type Publishers struct {
	Publishers []Publisher
}
