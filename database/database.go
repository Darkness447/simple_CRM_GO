package database

import (
	"fmt"
	"simple_crm_go/lead"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// capital is important
var DbConn *gorm.DB

func Initialize() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		panic(err)
	}

	fmt.Println("Connection to a database")
	db.AutoMigrate(&lead.Lead{})

	DbConn = db
}
