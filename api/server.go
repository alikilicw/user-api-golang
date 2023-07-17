package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "user-api/db/sqlc"
	"user-api/token"
	"user-api/util"
)

type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}
	router := gin.Default()

	server.router = router

	return server, nil
}
