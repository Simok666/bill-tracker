package services

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/dhani/bill-tracker-backend/internal/models"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

// CreateCategoryInput holds data for creating a category
type CreateCategoryInput struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	Icon        *string `json:"icon"`
	Color       *string `json:"color"`
}

// UpdateCategoryInput holds data for updating a category
type UpdateCategoryInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Icon        *string `json:"icon"`
	Color       *string `json:"color"`
}

// List retrieves all categories for a company
func (s *CategoryService) List(companyID uuid.UUID) ([]models.Category, error) {
	var categories []models.Category
	err := s.db.Where("company_id = ?", companyID).Order("name ASC").Find(&categories).Error
	return categories, err
}

// GetByID retrieves a category by ID
func (s *CategoryService) GetByID(companyID, categoryID uuid.UUID) (*models.Category, error) {
	var category models.Category
	err := s.db.Where("company_id = ? AND id = ?", companyID, categoryID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// Create creates a new category
func (s *CategoryService) Create(companyID uuid.UUID, input CreateCategoryInput) (*models.Category, error) {
	category := models.Category{
		CompanyID:   companyID,
		Name:        input.Name,
		Description: input.Description,
		Icon:        input.Icon,
		Color:       input.Color,
	}

	if err := s.db.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// Update updates an existing category
func (s *CategoryService) Update(companyID, categoryID uuid.UUID, input UpdateCategoryInput) (*models.Category, error) {
	var category models.Category
	if err := s.db.Where("company_id = ? AND id = ?", companyID, categoryID).First(&category).Error; err != nil {
		return nil, errors.New("category not found")
	}

	updates := make(map[string]interface{})
	if input.Name != nil {
		updates["name"] = *input.Name
	}
	if input.Description != nil {
		updates["description"] = *input.Description
	}
	if input.Icon != nil {
		updates["icon"] = *input.Icon
	}
	if input.Color != nil {
		updates["color"] = *input.Color
	}

	if err := s.db.Model(&category).Updates(updates).Error; err != nil {
		return nil, err
	}

	return s.GetByID(companyID, categoryID)
}

// Delete deletes a category
func (s *CategoryService) Delete(companyID, categoryID uuid.UUID) error {
	result := s.db.Where("company_id = ? AND id = ?", companyID, categoryID).Delete(&models.Category{})
	if result.RowsAffected == 0 {
		return errors.New("category not found")
	}
	return result.Error
}
