package server

import (
	"log"
	"mail-sender/config"
	"mail-sender/server/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() *Server {
	port, err := config.GetConfig("port")

	if err != nil || port == "" {
		port = "8080"
		log.Println("Server port wasn't specified, service will be started in default port: " + port)
	}

	return &Server{
		port: port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {

	router := routes.ConfigRoutes(s.server)

	log.Printf("Server is running on port %s", s.port)

	log.Fatal(router.Run(":" + s.port));
}