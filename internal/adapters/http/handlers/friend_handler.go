package handlers

import (
	"github.com/Reza-Rayan/internal/applications"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FriendHandler struct {
	friendUC *applications.FriendUsecase
}

func NewFriendHandler(friendUC *applications.FriendUsecase) *FriendHandler {
	return &FriendHandler{friendUC: friendUC}
}

// GetFriendsHandler -> GET method
func (h *FriendHandler) GetFriendsHandler(ctx *gin.Context) {
	userIDParam := ctx.Param("id")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	friends, err := h.friendUC.GetFriends(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"friends": friends})
}

func (h *FriendHandler) AddFriendHandler(ctx *gin.Context) {
	userIDParam := ctx.Param("id")
	friendIDParam := ctx.Param("friendId")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	friendID, err := strconv.ParseUint(friendIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid friend id"})
		return
	}

	err = h.friendUC.AddFriend(uint(userID), uint(friendID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "new friend added", "friend": friendID})
}
