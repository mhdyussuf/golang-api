package app

import (
	"log"
	"os"

	"golang-api/app/controllers"

	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	prod := os.Getenv("APP_ENV")
	if prod == "" {

		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error on Loading .env file")
		}
		appConfig.AppName = getEnv("APP_NAME", "IniAppCoba")
		appConfig.AppEnv = getEnv("APP_ENV", "develop")
		appConfig.AppPort = getEnv("APP_PORT", "9989")

		dbConfig.DBHost = getEnv("DB_HOST", "localhost")
		dbConfig.DBUser = getEnv("DB_USER", "user")
		dbConfig.DBPassword = getEnv("DB_PASSWORD", "password")
		dbConfig.DBName = getEnv("DB_NAME", "dbname")
		dbConfig.DBPort = getEnv("DB_PORT", "5432")
	} else {

		appConfig.AppName = os.Getenv("APP_NAME")
		appConfig.AppEnv = os.Getenv("APP_ENV")
		appConfig.AppPort = os.Getenv("APP_PORT")

		dbConfig.DBHost = os.Getenv("DB_HOST")
		dbConfig.DBUser = os.Getenv("DB_USER")
		dbConfig.DBPassword = os.Getenv("DB_PASSWORD")
		dbConfig.DBName = os.Getenv("DB_NAME")
		dbConfig.DBPort = os.Getenv("DB_PORT")

	}

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)

}
