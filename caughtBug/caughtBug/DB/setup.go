package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB
var server = "caughtbugdev.database.windows.net"
var port = 1433
var user = "caughtbugsql@caughtbugdev"
var password = "RogStrix@1080"
var database = "caughtbugsqldev"

func DBSetup() (*sql.DB, error) {
	// Build connection string
	var db *sql.DB
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		return db, err
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return db, err
	}
	log.Println("SQL Server Connected!")
	return db, nil
}
