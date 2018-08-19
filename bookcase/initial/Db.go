package initial

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SQL interface {
	Open(string, string) (disconnect func())
	Begin() error
	Close() error
}

type SQLInstance struct {
	DB *sql.DB
	Tx *sql.Tx
}

//// Open returns a DB reference for a data source.
func (s *SQLInstance) Open(driverName, dataSourceName string) (disconnect func()) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		fmt.Printf(err.Error())
		return func() {
			fmt.Println("something error occurs when try to establish the connection, so nothing will do to destroy the connection")
		}
	}

	s.DB = db

	return func() {
		fmt.Println("disconnect the connection")
		db.Close()
	}
}

// Begin starts an returns a new transaction.
func (s *SQLInstance) Begin() error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	s.Tx = tx
	return nil
}

// Begin starts an returns a new transaction.
func (s *SQLInstance) Commit() error {
	err := s.Tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLInstance) Close() error {
	log.Printf("close the db at %v", time.Now())
	return s.DB.Close()
}
