package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DBConnect *gorm.DB

func InitDB() {
	username := "root"
	password := ""
	host := "127.0.0.1"
	port := 3306
	dbname := "gotest"
	timeout := "10s"

	var mysqlLogger logger.Interface = logger.Default.LogMode(logger.Info)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&timeout=%s",
		username, password, host, port, dbname, timeout)

	var err error
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
	DBConnect.AutoMigrate(&Member{})
}

type Member struct {
	ID        uint      `gorm:"column:ID;primary_key"`
	Name      string    `gorm:"column:Name;type:varchar(60);not null"`
	Email     string    `gorm:"column:Email;type:varchar(100);unique_index;not null"`
	Password  string    `gorm:"column:Password;type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"column:CreatedAt;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt;type:datetime;default:CURRENT_TIMESTAMP"`
}
