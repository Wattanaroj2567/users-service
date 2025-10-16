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
// UserRepository defines the interface for user data operations.
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uint) (*models.User, error)
	FindByEmailOrUsername(ctx context.Context, identifier string) (*models.User, error)
=======
// UserRepository defines persistence operations for the users table.
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uint) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
}

<<<<<<< HEAD
// userRepository is the GORM implementation of the UserRepository interface.
=======
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
type userRepository struct {
	db *gorm.DB
}

<<<<<<< HEAD
// NewUserRepository creates a new GORM-backed UserRepository.
=======
// NewUserRepository wires a GORM-backed user repository.
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

<<<<<<< HEAD
// Create inserts a new user record into the database.
func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindByID retrieves a user by their primary key ID.
func (r *userRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmailOrUsername retrieves a user by their email or username.
func (r *userRepository) FindByEmailOrUsername(ctx context.Context, identifier string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil, nil if not found to distinguish from other errors
		}
		return nil, err
	}
	return &user, nil
}

// Update saves changes to an existing user record.
func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete removes a user record from the database.
func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}
=======
func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	// TODO: persist user record using GORM with proper error handling
	return nil
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	// TODO: fetch a user by primary key
	return nil, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	// TODO: fetch a user by email
	return nil, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	// TODO: fetch a user by username
	return nil, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	// TODO: update user record fields
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	// TODO: soft-delete or hard-delete user record
	return nil
}
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
