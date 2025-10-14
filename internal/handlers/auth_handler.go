package handlers

import (
	"net/http"

	"github.com/gamegear/users-service/internal/models"
	"github.com/gamegear/users-service/internal/services"
	"github.com/gin-gonic/gin"
)

// compile-time references to request DTOs avoid unused import warnings while logic is TODO.
var (
	_ models.RegisterRequest
	_ models.LoginRequest
	_ models.ForgotPasswordRequest
	_ models.ResetPasswordRequest
)

// AuthHandler wires HTTP endpoints to the AuthService.
type AuthHandler struct {
	authService services.AuthService
}

// NewAuthHandler constructs AuthHandler.
func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register handles POST /api/auth/register.
func (h *AuthHandler) Register(c *gin.Context) {
	// TODO: bind models.RegisterRequest, invoke service.Register, handle response/errors
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement register handler"})
}

// Login handles POST /api/auth/login.
func (h *AuthHandler) Login(c *gin.Context) {
	// TODO: bind models.LoginRequest, invoke service.Login, respond with token & user
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement login handler"})
}

// Logout handles POST /api/auth/logout.
func (h *AuthHandler) Logout(c *gin.Context) {
	// TODO: extract token from header, invoke service.Logout
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement logout handler"})
}

// ForgotPassword handles POST /api/auth/forgot-password.
func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	// TODO: bind models.ForgotPasswordRequest, invoke service.ForgotPassword
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement forgot-password handler"})
}

// ResetPassword handles POST /api/auth/reset-password.
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	// TODO: bind models.ResetPasswordRequest, invoke service.ResetPassword
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement reset-password handler"})
}

// TODO: consider returning typed responses for success/error once service layer is ready.
