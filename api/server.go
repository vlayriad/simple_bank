package api

import (
	"github.com/gin-gonic/gin"
	database "kaffein/simple_bank/database/sqlc"
)

type Server struct {
	store  *database.Store
	router *gin.Engine
}

func NewServer(s *database.Store) *Server {
	server := &Server{
		store:  s,
		router: gin.Default(),
	}

	server.router.POST("/accounts", server.createAccount)
	server.router.GET("/accounts/:id", server.getAccount)
	server.router.GET("/accounts", server.getListAccount)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
