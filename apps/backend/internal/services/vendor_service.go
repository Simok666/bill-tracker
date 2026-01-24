package services

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/dhani/bill-tracker-backend/internal/models"
)

type VendorService struct {
	db *gorm.DB
}

func NewVendorService(db *gorm.DB) *VendorService {
	return &VendorService{db: db}
}

// CreateVendorInput holds data for creating a vendor
type CreateVendorInput struct {
	Name         string  `json:"name" binding:"required"`
	LogoURL      *string `json:"logo_url"`
	ContactEmail *string `json:"contact_email"`
	Website      *string `json:"website"`
	ContactInfo  *string `json:"contact_info"`
	Address      *string `json:"address"`
	Location     *string `json:"location"`
}

// UpdateVendorInput holds data for updating a vendor
type UpdateVendorInput struct {
	Name         *string `json:"name"`
	LogoURL      *string `json:"logo_url"`
	ContactEmail *string `json:"contact_email"`
	Website      *string `json:"website"`
	ContactInfo  *string `json:"contact_info"`
	Address      *string `json:"address"`
	Location     *string `json:"location"`
}

// List retrieves all vendors for a company
func (s *VendorService) List(companyID uuid.UUID) ([]models.Vendor, error) {
	var vendors []models.Vendor
	err := s.db.Where("company_id = ?", companyID).Order("name ASC").Find(&vendors).Error
	return vendors, err
}

// GetByID retrieves a vendor by ID
func (s *VendorService) GetByID(companyID, vendorID uuid.UUID) (*models.Vendor, error) {
	var vendor models.Vendor
	err := s.db.Where("company_id = ? AND id = ?", companyID, vendorID).First(&vendor).Error
	if err != nil {
		return nil, err
	}
	return &vendor, nil
}

// Create creates a new vendor
func (s *VendorService) Create(companyID uuid.UUID, input CreateVendorInput) (*models.Vendor, error) {
	vendor := models.Vendor{
		CompanyID:    companyID,
		Name:         input.Name,
		LogoURL:      input.LogoURL,
		ContactEmail: input.ContactEmail,
		Website:      input.Website,
		ContactInfo:  input.ContactInfo,
		Address:      input.Address,
		Location:     input.Location,
	}

	if err := s.db.Create(&vendor).Error; err != nil {
		return nil, err
	}
	return &vendor, nil
}

// Update updates an existing vendor
func (s *VendorService) Update(companyID, vendorID uuid.UUID, input UpdateVendorInput) (*models.Vendor, error) {
	var vendor models.Vendor
	if err := s.db.Where("company_id = ? AND id = ?", companyID, vendorID).First(&vendor).Error; err != nil {
		return nil, errors.New("vendor not found")
	}

	updates := make(map[string]interface{})
	if input.Name != nil {
		updates["name"] = *input.Name
	}
	if input.LogoURL != nil {
		updates["logo_url"] = *input.LogoURL
	}
	if input.ContactEmail != nil {
		updates["contact_email"] = *input.ContactEmail
	}
	if input.Website != nil {
		updates["website"] = *input.Website
	}
	if input.ContactInfo != nil {
		updates["contact_info"] = *input.ContactInfo
	}
	if input.Address != nil {
		updates["address"] = *input.Address
	}
	if input.Location != nil {
		updates["location"] = *input.Location
	}

	if err := s.db.Model(&vendor).Updates(updates).Error; err != nil {
		return nil, err
	}

	return s.GetByID(companyID, vendorID)
}

// Delete deletes a vendor
func (s *VendorService) Delete(companyID, vendorID uuid.UUID) error {
	result := s.db.Where("company_id = ? AND id = ?", companyID, vendorID).Delete(&models.Vendor{})
	if result.RowsAffected == 0 {
		return errors.New("vendor not found")
	}
	return result.Error
}
