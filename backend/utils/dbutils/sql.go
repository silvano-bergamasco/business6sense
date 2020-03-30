package dbutils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (*sql.DB, error) {
	dbDriver := os.Getenv("MYSQL_DRIVER")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	bdurl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDriver, bdurl)

	if err != nil {
		return db, err
	}

	return db, nil
}

/*
var nt mysql.NullTime
err := db.QueryRow("SELECT time FROM foo WHERE id = ?", id).Scan(&nt)

if nt.Valid {
   // use nt.Time
} else {
   // NULL value
}
*/
