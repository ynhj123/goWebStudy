package mapper

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dsn   = "root:root123@tcp(127.0.0.1:3306)/u_chat"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
)
