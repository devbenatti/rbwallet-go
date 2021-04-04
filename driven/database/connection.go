package database

import (
	"database/sql"

	"github.com/devbenatti/rbwallet-go/config"
	"github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	config := config.Load()

	dbConfig := mysql.NewConfig()
	dbConfig.User = config.DBUser
	dbConfig.Passwd = config.DBPasswd
	dbConfig.Addr = config.DBAddr
	dbConfig.DBName = config.DBName
	dbConfig.Net = "tcp"

	db, err := sql.Open("mysql", dbConfig.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	return db
}
