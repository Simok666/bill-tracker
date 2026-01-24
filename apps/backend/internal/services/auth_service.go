package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/dhani/bill-tracker-backend/internal/middleware"
	"github.com/dhani/bill-tracker-backend/internal/models"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

// RegisterInput holds registration data
type RegisterInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	CompanyName string `json:"company_name" binding:"required"`
}

// LoginInput holds login data
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse is returned after successful auth
type AuthResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

// Register creates a new user and company
func (s *AuthService) Register(input RegisterInput) (*AuthResponse, error) {
	// Check if email already exists
	var existingUser models.User
	if err := s.db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create company
	company := models.Company{
		Name: input.CompanyName,
	}
	if err := s.db.Create(&company).Error; err != nil {
		return nil, errors.New("failed to create company")
	}

	// Create user
	user := models.User{
		CompanyID:    company.ID,
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
		Role:         models.RoleAdmin, // First user is admin
	}
	if err := s.db.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	// Generate token
	token, err := middleware.GenerateToken(&user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &AuthResponse{
		Token: token,
		User:  &user,
	}, nil
}

// Login authenticates a user
func (s *AuthService) Login(input LoginInput) (*AuthResponse, error) {
	var user models.User
	if err := s.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate token
	token, err := middleware.GenerateToken(&user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// Create session
	session := models.Session{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	s.db.Create(&session)

	return &AuthResponse{
		Token: token,
		User:  &user,
	}, nil
}

// Logout invalidates a session
func (s *AuthService) Logout(token string) error {
	return s.db.Where("token = ?", token).Delete(&models.Session{}).Error
}

// GetUserByID retrieves a user by ID
func (s *AuthService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
