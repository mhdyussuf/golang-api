package handlers

import (
	"net/http"
	"strconv"

	"golang-api/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MessageHandler struct {
	DB *gorm.DB
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	conversationId, err := strconv.Atoi(c.Param("conversation_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "err.Error()"})
		return
	}

	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check jika Conversation tidak ada
	var conversation models.Conversation
	if err := h.DB.First(&conversation, conversationId).Error; err != nil {

		// Auto membuat data untuk conversation ID baru
		conversation.IdConversation = uint(conversationId)
		if err := h.DB.Create(&conversation).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	message.ConversationId = uint(conversationId)
	if err := h.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.MessageResponse{
		IdMessage:      message.IdMessage,
		ConversationId: message.ConversationId,
		UserId:         message.UserId,
		Content:        message.Content,
		CreatedAt:      message.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)

}

func (h *MessageHandler) GetMessagesByConversationID(c *gin.Context) {
	conversationID, err := strconv.Atoi(c.Param("conversation_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var messages []models.Message
	if err := h.DB.Where("conversation_id = ?", conversationID).Find(&messages).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	var response []gin.H
	for _, message := range messages {
		response = append(response, gin.H{
			"id_message":      message.IdMessage,
			"conversation_id": message.ConversationId,
			"user_id":         message.UserId,
			"content":         message.Content,
			"sent_at":         message.CreatedAt,
		})
	}
	c.JSON(http.StatusOK, response)

}
