package DB

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectDatabase struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (dns *ConnectDatabase) dnsToString() (string, error) {
	if dns == nil {
		return "", fmt.Errorf("%s", "Connection data is null")
	}
	rpl := ""
	rpl += "host=" + dns.Host
	rpl += " port=" + dns.Port
	rpl += " user=" + dns.User
	rpl += " password=" + dns.Password
	rpl += " dbname=" + dns.DBName
	return rpl, nil
}
func (dns *ConnectDatabase) ConnectToDatabase() (*gorm.DB, error) {

	dnsString, err := dns.dnsToString()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dnsString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
