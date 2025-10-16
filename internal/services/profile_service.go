package services

import (
	"context"
<<<<<<< HEAD
	"errors"
	"fmt"

	"github.com/gamegear/users-service/internal/models"
	"github.com/gamegear/users-service/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrOldPasswordIncorrect = errors.New("old password is not correct")
	ErrPasswordRequired     = errors.New("password confirmation is required for this action")
)

// ProfileService defines the interface for profile management business logic.
=======

	"github.com/gamegear/users-service/internal/models"
	"github.com/gamegear/users-service/internal/repositories"
)

// ProfileService encapsulates member profile use cases.
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
type ProfileService interface {
	GetProfile(ctx context.Context, userID uint) (*models.ProfileResponse, error)
	UpdateProfile(ctx context.Context, userID uint, req models.UpdateProfileRequest) (*models.ProfileResponse, error)
}

<<<<<<< HEAD
// profileService is the implementation of the ProfileService interface.
=======
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
type profileService struct {
	userRepo repositories.UserRepository
}

<<<<<<< HEAD
// NewProfileService creates a new ProfileService.
=======
// NewProfileService constructs a profile service instance.
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
func NewProfileService(userRepo repositories.UserRepository) ProfileService {
	return &profileService{userRepo: userRepo}
}

<<<<<<< HEAD
// GetProfile retrieves a user's profile information.
func (s *profileService) GetProfile(ctx context.Context, userID uint) (*models.ProfileResponse, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	response := &models.ProfileResponse{
		ID:           user.ID,
		Username:     user.Username,
		DisplayName:  user.DisplayName,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
	}

	return response, nil
}

// UpdateProfile orchestrates the logic for updating profile info, changing the password, or deleting the account.
func (s *profileService) UpdateProfile(ctx context.Context, userID uint, req models.UpdateProfileRequest) (*models.ProfileResponse, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Case 1: Delete account
	if req.DeleteAccountFlag {
		if req.Password == "" {
			return nil, ErrPasswordRequired
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			return nil, ErrInvalidCredentials
		}
		if err := s.userRepo.Delete(ctx, userID); err != nil {
			return nil, fmt.Errorf("failed to delete account: %w", err)
		}
		return nil, nil
	}

	// Case 2: Change password
	if req.NewPassword != "" {
		if req.OldPassword == "" || req.ConfirmPassword == "" {
			return nil, errors.New("old password and confirmation are required to set a new password")
		}
		if req.NewPassword != req.ConfirmPassword {
			return nil, ErrPasswordMismatch
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.OldPassword)); err != nil {
			return nil, ErrOldPasswordIncorrect
		}
		newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("failed to hash new password: %w", err)
		}
		user.PasswordHash = string(newHashedPassword)
	}

	// Case 3: Update profile info
	if req.DisplayName != "" {
		user.DisplayName = req.DisplayName
	}
	if req.ProfileImage != "" {
		user.ProfileImage = req.ProfileImage
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update profile: %w", err)
	}

	return s.GetProfile(ctx, userID)
}
=======
func (s *profileService) GetProfile(ctx context.Context, userID uint) (*models.ProfileResponse, error) {
	// TODO: implement profile lookup and mapping
	return nil, nil
}

func (s *profileService) UpdateProfile(ctx context.Context, userID uint, req models.UpdateProfileRequest) (*models.ProfileResponse, error) {
	// TODO: implement profile update including password change and delete-account flag
	return nil, nil
}
>>>>>>> ed92ccd7167a49a8a8cf46a13d425b1d5fd62b92
