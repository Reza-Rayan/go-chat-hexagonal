package handlers

import (
	"github.com/Reza-Rayan/internal/applications"
	"github.com/Reza-Rayan/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

var cfg = config.LoadConfig()

type UserHandler struct {
	userUC *applications.UserUsecase
}

func NewUserHandler(userUC *applications.UserUsecase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

// SearchUsersHandler -> GET method
func (h *UserHandler) SearchUsersHandler(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "query parameter is required"})
		return
	}

	// pagination params
	limitStr := ctx.DefaultQuery("limit", "10")
	offsetStr := ctx.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	users, err := h.userUC.SearchUsers(query, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

// UpdateUserHandler -> PUT method
func (h *UserHandler) UpdateUserHandler(ctx *gin.Context) {
	userIDInterface, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	userID := userIDInterface.(uint)

	username := ctx.PostForm("username")
	email := ctx.PostForm("email")
	phone := ctx.PostForm("phone")

	var avatarURL string
	file, err := ctx.FormFile("avatar")
	if err == nil {
		filename := filepath.Base(file.Filename)
		savePath := filepath.Join(cfg.Server.UploadPath, filename)

		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			ctx.JSON(500, gin.H{"error": "failed to upload avatar"})
			return
		}

		avatarURL = cfg.Server.Addr + cfg.Server.UploadPath + filename
	}

	updatedUser, err := h.userUC.UpdateUser(userID, username, email, phone, avatarURL)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
		"user":    updatedUser,
	})
}
