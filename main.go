package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"user-api/api"
	"user-api/util"
)

func main() {
	config, err := util.LoadConfig(".")
	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db : ", err)
	}

	api.StartApi(conn)
}
