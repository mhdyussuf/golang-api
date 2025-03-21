package controllers

import (
	"golang-api/app/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// "github.com/gofiber/swagger"
	// _ "golang-api/docs"
)

func (server *Server) initializeRoutes() {
	server.Router = gin.Default()

	userHandler := handlers.UserHandler{DB: server.DB}
	messageHandler := handlers.MessageHandler{DB: server.DB}

	// User routes
	server.Router.POST("/users", userHandler.CreateUser)
	server.Router.GET("/users/:id", userHandler.GetUserByID)

	// Message routes
	server.Router.POST("/conversations/:conversation_id/messages", messageHandler.CreateMessage)
	server.Router.GET("/conversations/:conversation_id/messages", messageHandler.GetMessagesByConversationID)

	// Swagger
	// server.Router.Get("/swagger/*", swagger.HandlerDefault)
	server.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// log.Println("Server started at :8080")
	// server.Router.Run(":8080")
}
