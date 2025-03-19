package controllers

import (
	"fmt"
	"golang-api/app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}
type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Wellcome to " + appConfig.AppName)

	server.initializeDB(dbConfig)
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) initializeDB(dbConfig DBConfig) {

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBName)
	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed on connecting to the database server")
	}

	// untuk migrate database nya
	for _, model := range models.RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database Migrate Success")
}
