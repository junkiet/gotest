package main

import (
	"fmt"
	"goapi/db"
)

func init() {
	db.InitDB()

	// check the connection
	sqlDB, err := db.DBConnect.DB()
	if err != nil {
		panic("failed to get database")
	}
	// check if db connection is open
	err = sqlDB.Ping()
	if err != nil {
		panic("failed to ping database")
	}
	fmt.Println("Connected to the database!")
}

func main() {
}
