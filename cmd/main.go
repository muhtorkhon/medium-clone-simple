package main

import (
	"fmt"
	"log"

	"github.com/GofurovMuxtorxon/medium-clone-simple/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)



func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)
	psqConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("Filed to connect to postgres: %v", err)
	}

	log.Println("Postgres connection successfully done!")
	_ = psqConn
}