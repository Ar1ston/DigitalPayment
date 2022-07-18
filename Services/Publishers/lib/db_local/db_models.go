package db_local

type Publisher struct {
	Id          int64 `gorm:"primary key"`
	Name        string
	Description string
}
type Publishers struct {
	Publishers []Publisher
}
