package api

import (
	"github.com/gin-gonic/gin"
	database "github.com/terajari/idompet/database/sqlc"
)

type Server struct {
	Store  *database.Store
	Router *gin.Engine
}

func NewServer(db *database.Store) *Server {
	router := gin.Default()
	server := &Server{
		Store:  db,
		Router: router,
	}

	server.Router.POST("/akun", server.createAkun)
	server.Router.GET("/akun/:id", server.getAkun)
	server.Router.GET("/akun", server.listAkun)

	return server
}

func (server *Server) Start() {
	server.Router.Run()
}

func ErrResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
