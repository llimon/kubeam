package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*GetDatabaseConnection opens and pings a new database connection for test*/
func GetDatabaseConnection() *sql.DB {
	databaseIP, err := config.GetString("database/host", "mariadb")
	databasePort, err := config.GetInt("database/port", 3306)
	databasePassword, err := config.GetString("database/password", "")
	databaseType, err := config.GetString("database/type", "mysql")
	databaseName, err := config.GetString("database/name", "kubeam")
	databaseUser, err := config.GetString("database/user", "root")

	databaseConn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", databaseUser,
		databasePassword, databaseIP, databasePort, databaseName)

	db, err := sql.Open(databaseType, databaseConn)
	ErrorHandler(err)

	err = db.Ping()
	ErrorHandler(err)

	return db
}
