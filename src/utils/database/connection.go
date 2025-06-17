package database

import (
	config "digital-wallet/src/config"
	"fmt"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct{}

func (conn *Connection) Open() (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.NewDatabase().User,
		config.NewDatabase().Password,
		config.NewDatabase().Host,
		config.NewDatabase().Port,
		config.NewDatabase().Name,
	)

	// Open database connection
	// The database is called test
	db, err := sql.Open("mysql", connectionString)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(config.NewDatabase().MaxIdleConns)
	db.SetMaxOpenConns(config.NewDatabase().MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(config.NewDatabase().ConnMaxLifetime) * time.Second)

	fmt.Println("DB OPENED")

	return db, nil
}

func (conn *Connection) Close(db *sql.DB) error {
	err := db.Close()

	if err != nil {
		return fmt.Errorf("error when close db: %w", err)
	}

	fmt.Println("DB CLOSED")

	return nil
}
