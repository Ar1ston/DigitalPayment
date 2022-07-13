package db_local

type Author struct {
	Id          int64 `gorm:"primary key"`
	FirstName   string
	LastName    string
	Description string
}
type Authors struct {
	Authors []Author
}
