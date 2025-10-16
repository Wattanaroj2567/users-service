package repositories

import (
	"context"
<<<<<<< HEAD
	"errors"
=======
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92

	"github.com/gamegear/users-service/internal/models"
	"gorm.io/gorm"
)

<<<<<<< HEAD
// PasswordResetRepository defines the interface for managing password reset tokens.
=======
// PasswordResetRepository manages password reset tokens lifecycle.
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
type PasswordResetRepository interface {
	Create(ctx context.Context, token *models.PasswordResetToken) error
	FindByToken(ctx context.Context, token string) (*models.PasswordResetToken, error)
	DeleteByToken(ctx context.Context, token string) error
	DeleteByUserID(ctx context.Context, userID uint) error
}

<<<<<<< HEAD
// passwordResetRepository is the GORM implementation of the PasswordResetRepository interface.
=======
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
type passwordResetRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
// NewPasswordResetRepository creates a new GORM-backed PasswordResetRepository.
=======
// NewPasswordResetRepository wires a GORM-backed password reset repository.
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
func NewPasswordResetRepository(db *gorm.DB) PasswordResetRepository {
	return &passwordResetRepository{db: db}
}

<<<<<<< HEAD
// Create inserts a new password reset token record, after deleting any existing ones for the user.
func (r *passwordResetRepository) Create(ctx context.Context, token *models.PasswordResetToken) error {
	if err := r.DeleteByUserID(ctx, token.UserID); err != nil {
		return err
	}
	return r.db.WithContext(ctx).Create(token).Error
}

// FindByToken retrieves a password reset token by its token string.
func (r *passwordResetRepository) FindByToken(ctx context.Context, token string) (*models.PasswordResetToken, error) {
	var resetToken models.PasswordResetToken
	err := r.db.WithContext(ctx).Where("token = ?", token).First(&resetToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("token not found or has expired")
		}
		return nil, err
	}
	return &resetToken, nil
}

// DeleteByToken removes a password reset token from the database using the token string.
func (r *passwordResetRepository) DeleteByToken(ctx context.Context, token string) error {
	return r.db.WithContext(ctx).Where("token = ?", token).Delete(&models.PasswordResetToken{}).Error
}

// DeleteByUserID removes all password reset tokens associated with a specific user ID.
func (r *passwordResetRepository) DeleteByUserID(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&models.PasswordResetToken{}).Error
}
=======
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
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
