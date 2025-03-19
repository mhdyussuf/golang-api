package handlers

import (
	"net/http"
	"strconv"

	"golang-api/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

// CreateUser godoc
//
//	@Summary		Create an user
//	@Description	Create User
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.UserCreate	true	"query params"
//	@Success		200		{object}	models.UserResponse
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := models.UserResponse{
		IdUser:    user.IdUser,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// ShowUser godoc
//
//	@Summary		Show an user
//	@Description	get string by ID
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.UserResponse
//	@Failure		404	{object}	models.ErrorResponse
//	@Router			/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	response := models.UserResponse{
		IdUser:    user.IdUser,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}
