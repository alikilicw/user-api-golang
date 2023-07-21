package api

import (
	"database/sql"
	"log"
	db "user-api/db/sqlc"
	"user-api/util"
)

func StartApi(conn *sql.DB) {
	store := db.NewStore(conn)

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	server, err := NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	server.UrlPatterns(server.router)

	if err := server.router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
