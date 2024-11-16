package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"sync"
)

var (
	once sync.Once
	db   *pgx.Conn
)

func Init() error {
	return connect()
}

func GetInstance() *pgx.Conn {
	return db
}

func connect() error {
	var err error

	once.Do(func() {

		db, err = pgx.Connect(context.Background(),
			fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
				os.Getenv("DB_USERNAME"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_DATABASE"),
			),
		)
	})

	return err
}

func Close() error {
	if db == nil {
		return nil
	}

	return db.Close(context.Background())
}
