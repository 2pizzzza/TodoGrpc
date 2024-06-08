package postgres

import (
	"ToDoAppGrpc/internal/config"
	"ToDoAppGrpc/internal/lib/logger/sl"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"os"
)

type DB struct {
	db *sql.DB
}

func New(con *config.Config) (*DB, error) {

	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		con.Host, con.Port, con.Username, con.Database, con.Password)

	connDb, err := sql.Open("postgres", conn)

	if err != nil {
		log.Printf("Error connection db: %s", err)
		return nil, err
	}
	log.Printf("Succses connect database")

	err = ExecuteSQLFile(connDb)
	if err != nil {
		slog.Error("Failed create db", sl.Err(err))
		panic("good bye")
	}

	return &DB{
		db: connDb,
	}, nil
}

func ExecuteSQLFile(db *sql.DB) error {
	content, err := os.ReadFile("storage/initbd.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(content))
	return err
}
