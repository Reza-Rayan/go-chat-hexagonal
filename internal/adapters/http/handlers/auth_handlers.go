package handlers

import (
	"github.com/Reza-Rayan/internal/applications"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	authUC *applications.AuthUsecase
}

func NewAuthHandler(authUC *applications.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUC: authUC}
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authUC.Signup(req.Username, req.Password, req.Email, req.Phone, req.Avatar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully", "user": user})
}
