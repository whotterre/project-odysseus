package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/whotterre/odysseus/src/internal/services"
)

type AuthHandler struct {
	authService services.AuthService
	jwtSecret   string
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

func NewAuthHandler(authService services.AuthService, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		jwtSecret:   jwtSecret,
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

	if h.jwtSecret == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret is not configured"})
		return
	}

	expiresAt := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"iat":     time.Now().Unix(),
		"exp":     expiresAt.Unix(),
	})

	signedToken, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sign token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":                         signedToken,
		"account_public_key":            user.AccountPublicKey,
		"encrypted_account_private_key": user.EncryptedAccountPrivateKey,
		"device_public_key":             user.DevicePublicKey,
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
