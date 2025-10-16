package services

import (
	"context"
<<<<<<< HEAD
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gamegear/users-service/internal/models"
	"github.com/gamegear/users-service/internal/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrUserAlreadyExists     = errors.New("user with this email or username already exists")
	ErrPasswordMismatch      = errors.New("passwords do not match")
	ErrTokenInvalidOrExpired = errors.New("token is invalid or has expired")
)

// AuthService defines the interface for authentication-related business logic.
=======

	"github.com/gamegear/users-service/internal/models"
	"github.com/gamegear/users-service/internal/repositories"
)

// AuthService encapsulates authentication workflows for users-service.
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
type AuthService interface {
	Register(ctx context.Context, req models.RegisterRequest) (*models.AuthResponse, error)
	Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error)
	Logout(ctx context.Context, token string) error
	ForgotPassword(ctx context.Context, req models.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req models.ResetPasswordRequest) error
}

<<<<<<< HEAD
// authService is the implementation of the AuthService interface.
type authService struct {
	userRepo          repositories.UserRepository
	passwordResetRepo repositories.PasswordResetRepository
	tokenService      TokenService
}

// NewAuthService creates a new AuthService.
func NewAuthService(
	userRepo repositories.UserRepository,
	passwordResetRepo repositories.PasswordResetRepository,
	tokenService TokenService,
=======
type authService struct {
	userRepo          repositories.UserRepository
	passwordResetRepo repositories.PasswordResetRepository
}

// NewAuthService constructs the auth service implementation.
func NewAuthService(
	userRepo repositories.UserRepository,
	passwordResetRepo repositories.PasswordResetRepository,
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
) AuthService {
	return &authService{
		userRepo:          userRepo,
		passwordResetRepo: passwordResetRepo,
<<<<<<< HEAD
		tokenService:      tokenService,
	}
}

// Register orchestrates the user registration process.
func (s *authService) Register(ctx context.Context, req models.RegisterRequest) (*models.AuthResponse, error) {
	if req.Password != req.ConfirmPassword {
		return nil, ErrPasswordMismatch
	}

	existingUser, err := s.userRepo.FindByEmailOrUsername(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("error checking for existing user: %w", err)
	}
	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("could not hash password: %w", err)
	}

	user := &models.User{
		Username:     req.Username,
		DisplayName:  req.DisplayName,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Role:         "member",
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("could not create user: %w", err)
	}

	token, err := s.tokenService.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("could not generate token: %w", err)
	}

	user.PasswordHash = ""
	return &models.AuthResponse{User: user, Token: token}, nil
}

// Login orchestrates the user login process.
func (s *authService) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	user, err := s.userRepo.FindByEmailOrUsername(ctx, req.Identifier)
	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := s.tokenService.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, fmt.Errorf("could not generate token: %w", err)
	}

	user.PasswordHash = ""
	return &models.AuthResponse{User: user, Token: token}, nil
}

// Logout contains logic for user logout (e.g., token blacklisting if needed).
func (s *authService) Logout(ctx context.Context, token string) error {
	log.Printf("User logged out. Token revocation (if implemented) would apply to: %s", token)
	return nil
}

// ForgotPassword orchestrates the password reset initiation process.
func (s *authService) ForgotPassword(ctx context.Context, req models.ForgotPasswordRequest) error {
	user, err := s.userRepo.FindByEmailOrUsername(ctx, req.Email)
	if err != nil {
		return fmt.Errorf("error finding user: %w", err)
	}
	if user == nil {
		log.Printf("Password reset requested for non-existent email: %s", req.Email)
		return nil
	}

	resetToken := &models.PasswordResetToken{
		UserID:    user.ID,
		Token:     uuid.New().String(),
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	if err := s.passwordResetRepo.Create(ctx, resetToken); err != nil {
		return fmt.Errorf("could not create password reset token: %w", err)
	}

	// --- START: โค้ดที่ปรับปรุงใหม่ ---
	log.Println("!!! COPY THE TOKEN BELOW FOR THE NEXT STEP !!!")
	log.Printf("RESET_TOKEN_IS:%s", resetToken.Token)
	log.Println("!!! END OF TOKEN !!!")
	// --- END: โค้ดที่ปรับปรุงใหม่ ---

	return nil
}

// ResetPassword orchestrates the final step of the password reset process.
func (s *authService) ResetPassword(ctx context.Context, req models.ResetPasswordRequest) error {
	if req.NewPassword != req.ConfirmPassword {
		return ErrPasswordMismatch
	}

	resetToken, err := s.passwordResetRepo.FindByToken(ctx, req.Token)
	if err != nil {
		return ErrTokenInvalidOrExpired
	}

	if time.Now().After(resetToken.ExpiresAt) {
		_ = s.passwordResetRepo.DeleteByToken(ctx, req.Token)
		return ErrTokenInvalidOrExpired
	}

	user, err := s.userRepo.FindByID(ctx, resetToken.UserID)
	if err != nil {
		return ErrUserNotFound
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("could not hash new password: %w", err)
	}

	user.PasswordHash = string(newHashedPassword)
	if err := s.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("could not update password: %w", err)
	}

	return s.passwordResetRepo.DeleteByToken(ctx, req.Token)
}
=======
	}
}

func (s *authService) Register(ctx context.Context, req models.RegisterRequest) (*models.AuthResponse, error) {
	// TODO: implement registration flow (validate, hash, persist, emit welcome)
	return nil, nil
}

func (s *authService) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	// TODO: implement login flow (lookup by identifier, verify password, issue JWT)
	return nil, nil
}

func (s *authService) Logout(ctx context.Context, token string) error {
	// TODO: implement logout (token revocation strategy / blacklist)
	return nil
}

func (s *authService) ForgotPassword(ctx context.Context, req models.ForgotPasswordRequest) error {
	// TODO: implement forgot-password (create token, send email via notifier)
	return nil
}

func (s *authService) ResetPassword(ctx context.Context, req models.ResetPasswordRequest) error {
	// TODO: implement password reset (validate token, hash new password, cleanup)
	return nil
}
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
