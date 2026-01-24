package services

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/dhani/bill-tracker-backend/internal/models"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// UpdateProfileInput holds profile update data
type UpdateProfileInput struct {
	Name      *string `json:"name"`
	AvatarURL *string `json:"avatar_url"`
}

// ChangePasswordInput holds password change data
type ChangePasswordInput struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
}

// GetProfile retrieves user profile with company info
func (s *UserService) GetProfile(userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := s.db.Preload("Company").First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateProfile updates user profile
func (s *UserService) UpdateProfile(userID uuid.UUID, input UpdateProfileInput) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	updates := make(map[string]interface{})
	if input.Name != nil {
		updates["name"] = *input.Name
	}
	if input.AvatarURL != nil {
		updates["avatar_url"] = *input.AvatarURL
	}

	if err := s.db.Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}

	return s.GetProfile(userID)
}

// ChangePassword changes user password
func (s *UserService) ChangePassword(userID uuid.UUID, input ChangePasswordInput) error {
	var user models.User
	if err := s.db.First(&user, "id = ?", userID).Error; err != nil {
		return errors.New("user not found")
	}

	// Verify current password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.CurrentPassword)); err != nil {
		return errors.New("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	return s.db.Model(&user).Update("password_hash", string(hashedPassword)).Error
}
