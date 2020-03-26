package dbutils

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (*sql.DB, error) {
	dbDriver := os.Getenv("MYSQL_DRIVER")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := s.Getenv("MYSQL_PASSWORD")
	dbName := s.Getenv("MYSQL_DATABASE")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		return db, err
	}

	return db, nil
}
