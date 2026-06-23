package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/whotterre/odysseus/src/internal/services"
)

type AuthHandler struct {
	authService services.AuthService
}

type SignupUserRequest struct {
	Email                      string `json:"email" binding:"required,email"`
	Password                   string `json:"password" binding:"required"`
	AccountPublicKey           string `json:"account_public_key" binding:"required"`
	EncryptedAccountPrivateKey string `json:"encrypted_account_private_key" binding:"required"`
	DevicePublicKey            string `json:"device_public_key" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) LoginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Valid email and password required"})
		return
	}

	user, err := h.authService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"account_public_key":            user.AccountPublicKey,
		"encrypted_account_private_key": user.EncryptedAccountPrivateKey,
		"device_public_key":            user.DevicePublicKey,
	})
}

func (h *AuthHandler) SignupUser(ctx *gin.Context) {
	var req SignupUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields or malformed JSON"})
		return
	}

	id, err := h.authService.RegisterUser(
		req.Email, 
		req.Password, 
		req.AccountPublicKey, 
		req.EncryptedAccountPrivateKey, 
		req.DevicePublicKey,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"device_id": id,
		"status":    "success",
	})
}