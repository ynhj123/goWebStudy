package mapper

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dsn   = "root:root123@tcp(127.0.0.1:3306)/u_chat?useUnicode=true&characterEncoding=utf8mb4&allowMultiQueries=true&rewriteBatchedStatements=true&parseTime=True"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
)
