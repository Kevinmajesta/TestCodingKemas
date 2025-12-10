package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/db_kemas?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
