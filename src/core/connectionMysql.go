package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql() (*gorm.DB, error) {
	var DNS = "root:admin@tcp(localhost:3306)/words?parseTime=true"
	
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	fmt.Println("Connected to database successfully")
	return db, nil
}
