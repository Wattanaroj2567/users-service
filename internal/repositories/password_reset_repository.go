package repositories

import (
	"context"

	"github.com/gamegear/users-service/internal/models"
	"gorm.io/gorm"
)

// PasswordResetRepository manages password reset tokens lifecycle.
type PasswordResetRepository interface {
	Create(ctx context.Context, token *models.PasswordResetToken) error
	FindByToken(ctx context.Context, token string) (*models.PasswordResetToken, error)
	DeleteByToken(ctx context.Context, token string) error
	DeleteByUserID(ctx context.Context, userID uint) error
}

type passwordResetRepository struct {
	db *gorm.DB
}

// NewPasswordResetRepository wires a GORM-backed password reset repository.
func NewPasswordResetRepository(db *gorm.DB) PasswordResetRepository {
	return &passwordResetRepository{db: db}
}

func (r *passwordResetRepository) Create(ctx context.Context, token *models.PasswordResetToken) error {
	// TODO: insert password reset token record
	return nil
}

func (r *passwordResetRepository) FindByToken(ctx context.Context, token string) (*models.PasswordResetToken, error) {
	// TODO: lookup reset token and associated user
	return nil, nil
}

func (r *passwordResetRepository) DeleteByToken(ctx context.Context, token string) error {
	// TODO: remove reset token by token string
	return nil
}

func (r *passwordResetRepository) DeleteByUserID(ctx context.Context, userID uint) error {
	// TODO: cleanup tokens for a specific user
	return nil
}
