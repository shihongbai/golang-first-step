package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接信息
const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "bsh820102"
	dbName   = "shbai_db_test"
)

func Connect() sql.DB {
	// 连接数据库
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName))
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	return *db
}
