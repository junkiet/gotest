package main

import (
	"encoding/json"
	"fmt"
	"goapi/model"
	"goapi/mysql"

	"gorm.io/gorm"
)

// set glbal for db connection
var DB *gorm.DB

func init() {
	mysql.Connect()

	// check the connection
	sqlDB, err := mysql.DBConnect.DB()
	if err != nil {
		panic("failed to get database")
	}
	// check if db connection is open
	err = sqlDB.Ping()
	if err != nil {
		panic("failed to ping database")
	}
	// Assign the global DB variable
	DB = mysql.DBConnect
	fmt.Println("Connected to the database!")
}

func main() {
	var members []model.Member
	// find all and order by ID in descending order
	ids := []uint{1, 1165}
	result := DB.Select("Name", "Email").Find(&members, ids)
	if result.Error != nil {
		fmt.Println("Error fetching:", result.Error)
		return
	}
	if result.RowsAffected == 0 {
		fmt.Println("No records found.")
		return
	}

	for _, m := range members {
		fmt.Println(m.Name, m.Email)
	}
	// convert to json
	jsonMember, err := json.MarshalIndent(members, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonMember))
}
