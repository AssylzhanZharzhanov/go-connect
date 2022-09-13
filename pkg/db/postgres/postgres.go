package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return db, nil
}
