package repositories

import (
	"context"

	"github.com/gamegear/users-service/internal/models"
	"gorm.io/gorm"
)

// UserRepository defines persistence operations for the users table.
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uint) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository wires a GORM-backed user repository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

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
