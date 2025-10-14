package services

import (
	"context"

	"github.com/gamegear/users-service/internal/models"
	"github.com/gamegear/users-service/internal/repositories"
)

// AuthService encapsulates authentication workflows for users-service.
type AuthService interface {
	Register(ctx context.Context, req models.RegisterRequest) (*models.AuthResponse, error)
	Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error)
	Logout(ctx context.Context, token string) error
	ForgotPassword(ctx context.Context, req models.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req models.ResetPasswordRequest) error
}

type authService struct {
	userRepo          repositories.UserRepository
	passwordResetRepo repositories.PasswordResetRepository
}

// NewAuthService constructs the auth service implementation.
func NewAuthService(
	userRepo repositories.UserRepository,
	passwordResetRepo repositories.PasswordResetRepository,
) AuthService {
	return &authService{
		userRepo:          userRepo,
		passwordResetRepo: passwordResetRepo,
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
