package handlers

import (
	"github.com/Reza-Rayan/internal/applications"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userUC *applications.UserUsecase
}

func NewUserHandler(userUC *applications.UserUsecase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

func (u *UserHandler) SearchUsersHandler(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Query parameter is missing"})
		return
	}
	users, err := u.userUC.SearchUsers(query)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"results:": users})
}
