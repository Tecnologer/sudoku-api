package data

import (
	"fmt"
	"log"

	"github.com/tecnologer/sudoku/clients/sudoku-api/auth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	initialMigration()
}

func GetDatabase() *gorm.DB {
	// dbUsername := secrets.GetKeyString("sqlite.password")
	// dbPassword := secrets.GetKeyString("sqlite.password")

	config := &gorm.Config{}
	// databaseurl := fmt.Sprintf("file:sudoku-api.s3db?_auth&_auth_user=%s&_auth_pass=%s", dbUsername, dbPassword)
	connection, err := gorm.Open(sqlite.Open("file:sudoku-api.s3db"), config)
	if err != nil {
		log.Fatalln("wrong database url")
	}

	sqldb, err := connection.DB()

	if err != nil {
		log.Fatalln("something wrong with DB")
	}

	err = sqldb.Ping()
	if err != nil {
		log.Fatal("database not connected")
	}

	fmt.Println("connected to database")
	return connection
}

func Closedatabase(connection *gorm.DB) {
	sqldb, _ := connection.DB()
	sqldb.Close()
}

func initialMigration() {
	connection := GetDatabase()
	defer Closedatabase(connection)
	connection.AutoMigrate(auth.User{})
}
