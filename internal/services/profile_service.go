package services

import (
	"context"

	"github.com/gamegear/users-service/internal/models"
	"github.com/gamegear/users-service/internal/repositories"
)

// ProfileService encapsulates member profile use cases.
type ProfileService interface {
	GetProfile(ctx context.Context, userID uint) (*models.ProfileResponse, error)
	UpdateProfile(ctx context.Context, userID uint, req models.UpdateProfileRequest) (*models.ProfileResponse, error)
}

type profileService struct {
	userRepo repositories.UserRepository
}

// NewProfileService constructs a profile service instance.
func NewProfileService(userRepo repositories.UserRepository) ProfileService {
	return &profileService{userRepo: userRepo}
}

func (s *profileService) GetProfile(ctx context.Context, userID uint) (*models.ProfileResponse, error) {
	// TODO: implement profile lookup and mapping
	return nil, nil
}

func (s *profileService) UpdateProfile(ctx context.Context, userID uint, req models.UpdateProfileRequest) (*models.ProfileResponse, error) {
	// TODO: implement profile update including password change and delete-account flag
	return nil, nil
}
