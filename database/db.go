package database

import (
	"fmt"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
type TaxiConn struct {
	Db *gorm.DB
}

// todo: move to config file, put here rfor now
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "Taxi"
)

func Connect() *gorm.DB {
	once.Do(func()	{
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
		var err error
		db, err = gorm.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
	})
	return db
}