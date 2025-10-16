package handlers

import (
<<<<<<< HEAD
	"errors"
=======
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
	"net/http"

	"github.com/gamegear/users-service/internal/models"
	"github.com/gamegear/users-service/internal/services"
	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
// AuthHandler handles all authentication-related HTTP requests.
=======
// compile-time references to request DTOs avoid unused import warnings while logic is TODO.
var (
	_ models.RegisterRequest
	_ models.LoginRequest
	_ models.ForgotPasswordRequest
	_ models.ResetPasswordRequest
)

// AuthHandler wires HTTP endpoints to the AuthService.
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
type AuthHandler struct {
	authService services.AuthService
}

<<<<<<< HEAD
// NewAuthHandler creates a new AuthHandler.
=======
// NewAuthHandler constructs AuthHandler.
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

<<<<<<< HEAD
// Register handles the user registration request (POST /api/auth/register).
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.authService.Register(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, services.ErrUserAlreadyExists) || errors.Is(err, services.ErrPasswordMismatch) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// Login handles the user login request (POST /api/auth/login).
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Logout handles the user logout request (POST /api/auth/logout).
func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// ForgotPassword handles the request to initiate a password reset (POST /api/auth/forgot-password).
func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var req models.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body, email is required"})
		return
	}

	if err := h.authService.ForgotPassword(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process forgot password request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "If an account with that email exists, a password reset link has been sent."})
}

// ResetPassword handles setting a new password using a reset token (POST /api/auth/reset-password).
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req models.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.authService.ResetPassword(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, services.ErrTokenInvalidOrExpired) || errors.Is(err, services.ErrPasswordMismatch) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password has been reset successfully."})
}
=======
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
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
