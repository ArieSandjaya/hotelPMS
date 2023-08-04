package config

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Db() *gorm.DB {
	configuration := GetConfig()
	dbUserName := configuration.DB_USERNAME
	dbPassword := configuration.DB_PASSWORD
	dbHost := configuration.DB_HOST
	dbPort := configuration.DB_PORT
	dbName := configuration.DB_NAME
	dbMaxIdleTime := configuration.DB_CONN_MAX_IDLE_TIME
	dbMaxIdleConn := configuration.DB_MAX_IDLE_CONNS
	dbMaxOpenConn := configuration.DB_MAX_OPEN_CONNS

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUserName, dbPassword, dbHost, dbPort, dbName)

	mySQLDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	mySQLDB.SetConnMaxIdleTime(dbMaxIdleTime)
	mySQLDB.SetMaxIdleConns(dbMaxIdleConn)
	mySQLDB.SetMaxOpenConns(dbMaxOpenConn)

	db, err := gorm.Open(mysql.New(mysql.Config{Conn: mySQLDB}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true}})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
