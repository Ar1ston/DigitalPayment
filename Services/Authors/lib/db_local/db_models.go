package db_local

type Author struct {
	Id          int64  `gorm:"primary key"`
	FirstName   string `gorm:"column:FirstName"`
	LastName    string `gorm:"column:LastName"`
	Description string `gorm:"column:Description"`
}
type Authors struct {
	Authors []Author
}
