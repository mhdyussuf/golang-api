package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-api/app/handlers"
	"golang-api/app/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTestDBMessage() *gorm.DB {
	dsn := "root:admin@tcp(localhost:3306)/db_golang_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Message{}, &models.Conversation{})
	return db
}

func SetUpRouterMessage() *gin.Engine {
	router := gin.Default()
	return router
}
func TestCreateMessage(t *testing.T) {
	db := setupTestDBMessage()
	messageHandler := handlers.MessageHandler{DB: db}

	router := SetUpRouterMessage()
	router.POST("/conversations/:conversation_id/messages", messageHandler.CreateMessage)

	message := models.Message{
		UserId:  1,
		Content: "Hai ",
	}

	jsonValue, _ := json.Marshal(message)
	req, _ := http.NewRequest("POST", "/conversations/1/messages", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetMessagesByConversationID(t *testing.T) {
	db := setupTestDBMessage()
	messageHandler := handlers.MessageHandler{DB: db}

	router := SetUpRouterMessage()
	router.GET("/conversations/:conversation_id/messages", messageHandler.GetMessagesByConversationID)

	user := models.User{
		Username: "johndoe11",
		Email:    "john11@example.com",
		Password: "password123",
	}
	db.Create(&user)

	conversation := models.Conversation{}
	db.Create(&conversation)

	// Create test messages
	message1 := models.Message{
		UserId:         user.IdUser,
		ConversationId: conversation.IdConversation,
		Content:        "Test message 1",
	}
	db.Create(&message1)

	message2 := models.Message{
		UserId:         user.IdUser,
		ConversationId: conversation.IdConversation,
		Content:        "Test message 2",
	}
	db.Create(&message2)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/conversations/%d/messages", conversation.IdConversation), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var messages []models.Message
	err := json.Unmarshal(w.Body.Bytes(), &messages)
	assert.Nil(t, err)

	// Check messages data
	assert.Equal(t, 2, len(messages))
	assert.Equal(t, message1.Content, messages[0].Content)
	assert.Equal(t, message2.Content, messages[1].Content)
}
