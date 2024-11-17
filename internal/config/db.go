package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func DBConnection() *pgx.Conn {
	dburl := os.Getenv("DBURL")

	conn, err := pgx.Connect(context.Background(), dburl)
	if err != nil {
		log.Fatal("failed to open db", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal("failed to ping db", err)
	}

	log.Println("db connected")
	return conn
}
