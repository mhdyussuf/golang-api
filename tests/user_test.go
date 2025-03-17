package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"golang-api/app/handlers"
	"golang-api/app/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn := "root:admin@tcp(localhost:3306)/db_golang_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Message{}, &models.Conversation{})
	return db
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB()
	userHandler := handlers.UserHandler{DB: db}

	router := SetUpRouter()
	router.POST("/users", userHandler.CreateUser)

	// generate unique email , unique usernam
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	uniqueUsername := fmt.Sprintf("johndoe_%s", timestamp)
	uniqueEmail := fmt.Sprintf("john_%s@example.com", timestamp)

	user := models.User{
		Username: uniqueUsername,
		Email:    uniqueEmail,
		Password: "password123",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestGetUserByID(t *testing.T) {
	db := setupTestDB()
	userHandler := handlers.UserHandler{DB: db}

	router := SetUpRouter()
	router.GET("/users/:id", userHandler.GetUserByID)

	// generate unique email , unique usernam
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	uniqueUsername := fmt.Sprintf("johndoe_%s", timestamp)
	uniqueEmail := fmt.Sprintf("john_%s@example.com", timestamp)

	user := models.User{
		Username: uniqueUsername,
		Email:    uniqueEmail,
		Password: "password123",
	}
	db.Create(&user)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", user.IdUser), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resultUser models.User
	json.Unmarshal(w.Body.Bytes(), &resultUser)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, resultUser)
}
