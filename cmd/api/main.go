package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gamegear/users-service/internal/handlers"
	"github.com/gamegear/users-service/internal/models"
	"github.com/gamegear/users-service/internal/repositories"
	"github.com/gamegear/users-service/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables so the service matches deployment configuration.
	if err := godotenv.Load(); err != nil {
		log.Println("warn: no .env file found, relying on environment variables")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is required")
	}

	// Establish database connection.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Apply schema migrations expected by README (users, password reset tokens).
	if err := db.AutoMigrate(&models.User{}, &models.PasswordResetToken{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Wire repositories.
	userRepo := repositories.NewUserRepository(db)
	passwordResetRepo := repositories.NewPasswordResetRepository(db)

	// Wire services.
	authService := services.NewAuthService(userRepo, passwordResetRepo)
	profileService := services.NewProfileService(userRepo)

	// Wire handlers.
	authHandler := handlers.NewAuthHandler(authService)
	profileHandler := handlers.NewProfileHandler(profileService)

	// Setup Gin router with standard middleware.
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// Health check as documented in README.
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Register versioned API routes.
	handlers.RegisterRoutes(router, authHandler, profileHandler)

	// Start HTTP server.
	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("users-service ready on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
