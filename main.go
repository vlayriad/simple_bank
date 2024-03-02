package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"kaffein/simple_bank/api"
	database "kaffein/simple_bank/database/sqlc"
	"kaffein/simple_bank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")

	connection, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connection to database ", err.Error())
	}
	store := database.NewStore(connection)
	server := api.NewServer(store)

	err = server.Start(fmt.Sprintf("%s:%s", config.DBHost, config.DBPort))
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
