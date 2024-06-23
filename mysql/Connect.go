package mysql

import (
	"fmt"

	"goapi/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Connect() {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := 3306
	dbname := "wm"
	timeout := "10s"

	var mysqlLogger logger.Interface = logger.Default.LogMode(logger.Info)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&timeout=%s",
		username, password, host, port, dbname, timeout)

	var err error
	var DBConnect *gorm.DB
	DBConnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "jk_",
			SingularTable: false,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	// sync the schema
	// when table exist, dont auto migrate the table
	DBConnect.AutoMigrate(&model.Member{})
}
