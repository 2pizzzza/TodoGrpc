package postgres

import (
	"ToDoAppGrpc/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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

	return &DB{
		db: connDb,
	}, nil
}
