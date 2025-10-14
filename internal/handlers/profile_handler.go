package handlers

import (
	"net/http"
	"strconv"

	"github.com/gamegear/users-service/internal/models"
	"github.com/gamegear/users-service/internal/services"
	"github.com/gin-gonic/gin"
)

// compile-time references to request DTOs avoid unused import warnings while logic is TODO.
var (
	_ models.ProfileResponse
	_ models.UpdateProfileRequest
)

// ProfileHandler exposes profile-related routes.
type ProfileHandler struct {
	profileService services.ProfileService
}

// NewProfileHandler constructs ProfileHandler.
func NewProfileHandler(profileService services.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

// GetProfile handles GET /api/user/profile.
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	// TODO: derive userID from authentication middleware / token claims instead of query parameter
	_, _ = strconv.Atoi(c.Query("user_id"))
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement get profile handler"})
}

// UpdateProfile handles PUT /api/user/profile.
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	// TODO: bind models.UpdateProfileRequest and invoke profileService.UpdateProfile
	c.JSON(http.StatusNotImplemented, gin.H{"message": "TODO: implement update profile handler"})
}

// NOTE: actual user identification will rely on middleware parsing JWTs once implemented.
