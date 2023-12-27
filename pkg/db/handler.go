package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBHandler struct {
	Conn *gorm.DB
}

func connect(host string, port string, user string, dbName, password string) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	conn, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func New(host string, port string, user string, dbName string, password string) (*DBHandler, error) {
	conn, err := connect(host, port, user, dbName, password)
	if err != nil {
		return nil, err
	}
	dbHandler := new(DBHandler)
	dbHandler.Conn = conn
	return dbHandler, nil
}
