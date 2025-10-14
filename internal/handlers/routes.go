package handlers

import "github.com/gin-gonic/gin"

// RegisterRoutes mounts all HTTP routes exposed by users-service following README specification.
func RegisterRoutes(router *gin.Engine, authHandler *AuthHandler, profileHandler *ProfileHandler) {
	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
		auth.POST("/forgot-password", authHandler.ForgotPassword)
		auth.POST("/reset-password", authHandler.ResetPassword)
	}

	user := api.Group("/user")
	{
		user.GET("/profile", profileHandler.GetProfile)
		user.PUT("/profile", profileHandler.UpdateProfile)
	}
}
