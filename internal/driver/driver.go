package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var databaseConnection = &DB{}

const maxOpenDatabaseConnections = 5
const maxIddleDatabaseConnections = 5
const maxDatabaseLifeTime = 5 * time.Minute

func ConnectPostgres(dataSourceName string) (*DB, error) {
	database, err := sql.Open("pgx", dataSourceName)

	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(maxOpenDatabaseConnections)
	database.SetMaxIdleConns(maxIddleDatabaseConnections)
	database.SetConnMaxLifetime(maxDatabaseLifeTime)

	err = testDB(database)

	if err != nil {
		return nil, err
	}

	databaseConnection.SQL = database

	return databaseConnection, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()

	if err != nil {
		return err
	}
	fmt.Println("*** Pinged database successfully! ***")
	return nil
}
