package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xhoang0509/simplebank/api"
	db "github.com/xhoang0509/simplebank/db/sqlc"
	"github.com/xhoang0509/simplebank/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config: ", err)
	}
	dbConnect, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	conn := db.NewStore(dbConnect)

	server := api.NewServer(conn)

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatalln("cannot start server: ", err)
	}

	log.Println("server is running on port: ", config.ServerAddress)
}
