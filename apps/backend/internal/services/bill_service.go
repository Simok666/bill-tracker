package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/dhani/bill-tracker-backend/internal/models"
	"github.com/dhani/bill-tracker-backend/internal/utils"
)

type BillService struct {
	db *gorm.DB
}

func NewBillService(db *gorm.DB) *BillService {
	return &BillService{db: db}
}

// BillFilters holds query filters
type BillFilters struct {
	Status string
	Search string
	utils.Pagination
}

// CreateBillInput holds data for creating a bill
type CreateBillInput struct {
	Title              string                    `json:"title" binding:"required"`
	VendorID           *uuid.UUID                `json:"vendor_id"`
	CategoryID         *uuid.UUID                `json:"category_id"`
	InvoiceNumber      *string                   `json:"invoice_number"`
	Amount             decimal.Decimal           `json:"amount" binding:"required"`
	Currency           string                    `json:"currency"`
	DueDate            time.Time                 `json:"due_date" binding:"required"`
	IsRecurring        bool                      `json:"is_recurring"`
	RecurringFrequency *models.RecurringFrequency `json:"recurring_frequency"`
	RecurringDay       *int                      `json:"recurring_day"`
	PaymentMethod      *string                   `json:"payment_method"`
	Notes              *string                   `json:"notes"`
	Status             models.BillStatus         `json:"status"`
}

// UpdateBillInput holds data for updating a bill
type UpdateBillInput struct {
	Title              *string                    `json:"title"`
	VendorID           *uuid.UUID                 `json:"vendor_id"`
	CategoryID         *uuid.UUID                 `json:"category_id"`
	InvoiceNumber      *string                    `json:"invoice_number"`
	Amount             *decimal.Decimal           `json:"amount"`
	Currency           *string                    `json:"currency"`
	DueDate            *time.Time                 `json:"due_date"`
	IsRecurring        *bool                      `json:"is_recurring"`
	RecurringFrequency *models.RecurringFrequency `json:"recurring_frequency"`
	RecurringDay       *int                       `json:"recurring_day"`
	PaymentMethod      *string                    `json:"payment_method"`
	Notes              *string                    `json:"notes"`
	Status             *models.BillStatus         `json:"status"`
}

// List retrieves bills with filters and pagination
func (s *BillService) List(companyID uuid.UUID, filters BillFilters) ([]models.Bill, int64, error) {
	var bills []models.Bill
	var total int64

	query := s.db.Model(&models.Bill{}).Where("company_id = ?", companyID)

	// Apply status filter
	if filters.Status != "" && filters.Status != "all" {
		query = query.Where("status = ?", filters.Status)
	}

	// Apply search
	if filters.Search != "" {
		searchTerm := "%" + filters.Search + "%"
		query = query.Where(
			"title ILIKE ? OR invoice_number ILIKE ?",
			searchTerm, searchTerm,
		)
	}

	// Get total count
	query.Count(&total)

	// Apply pagination
	err := query.
		Preload("Vendor").
		Preload("Category").
		Order(filters.Pagination.GetOrderClause()).
		Offset(filters.Pagination.GetOffset()).
		Limit(filters.Pagination.PageSize).
		Find(&bills).Error

	return bills, total, err
}

// GetByID retrieves a single bill by ID
func (s *BillService) GetByID(companyID, billID uuid.UUID) (*models.Bill, error) {
	var bill models.Bill
	err := s.db.
		Preload("Vendor").
		Preload("Category").
		Preload("Attachments").
		Preload("Activities", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC").Limit(10)
		}).
		Preload("Activities.User").
		Where("company_id = ? AND id = ?", companyID, billID).
		First(&bill).Error

	if err != nil {
		return nil, err
	}
	return &bill, nil
}

// Create creates a new bill
func (s *BillService) Create(companyID, userID uuid.UUID, input CreateBillInput) (*models.Bill, error) {
	status := input.Status
	if status == "" {
		status = models.StatusDraft
	}

	currency := input.Currency
	if currency == "" {
		currency = "USD"
	}

	bill := models.Bill{
		CompanyID:          companyID,
		UserID:             userID,
		VendorID:           input.VendorID,
		CategoryID:         input.CategoryID,
		Title:              input.Title,
		InvoiceNumber:      input.InvoiceNumber,
		Amount:             input.Amount,
		Currency:           currency,
		DueDate:            input.DueDate,
		Status:             status,
		IsRecurring:        input.IsRecurring,
		RecurringFrequency: input.RecurringFrequency,
		RecurringDay:       input.RecurringDay,
		PaymentMethod:      input.PaymentMethod,
		Notes:              input.Notes,
	}

	if err := s.db.Create(&bill).Error; err != nil {
		return nil, err
	}

	// Log activity
	s.logActivity(bill.ID, &userID, models.ActionCreated, "Bill created")

	return &bill, nil
}

// Update updates an existing bill
func (s *BillService) Update(companyID, billID uuid.UUID, userID uuid.UUID, input UpdateBillInput) (*models.Bill, error) {
	var bill models.Bill
	if err := s.db.Where("company_id = ? AND id = ?", companyID, billID).First(&bill).Error; err != nil {
		return nil, errors.New("bill not found")
	}

	updates := make(map[string]interface{})

	if input.Title != nil {
		updates["title"] = *input.Title
	}
	if input.VendorID != nil {
		updates["vendor_id"] = *input.VendorID
	}
	if input.CategoryID != nil {
		updates["category_id"] = *input.CategoryID
	}
	if input.InvoiceNumber != nil {
		updates["invoice_number"] = *input.InvoiceNumber
	}
	if input.Amount != nil {
		updates["amount"] = *input.Amount
	}
	if input.Currency != nil {
		updates["currency"] = *input.Currency
	}
	if input.DueDate != nil {
		updates["due_date"] = *input.DueDate
	}
	if input.IsRecurring != nil {
		updates["is_recurring"] = *input.IsRecurring
	}
	if input.RecurringFrequency != nil {
		updates["recurring_frequency"] = *input.RecurringFrequency
	}
	if input.RecurringDay != nil {
		updates["recurring_day"] = *input.RecurringDay
	}
	if input.PaymentMethod != nil {
		updates["payment_method"] = *input.PaymentMethod
	}
	if input.Notes != nil {
		updates["notes"] = *input.Notes
	}
	if input.Status != nil {
		updates["status"] = *input.Status
	}

	if err := s.db.Model(&bill).Updates(updates).Error; err != nil {
		return nil, err
	}

	// Log activity
	s.logActivity(bill.ID, &userID, models.ActionUpdated, "Bill updated")

	return s.GetByID(companyID, billID)
}

// Delete soft-deletes a bill
func (s *BillService) Delete(companyID, billID uuid.UUID) error {
	result := s.db.Where("company_id = ? AND id = ?", companyID, billID).Delete(&models.Bill{})
	if result.RowsAffected == 0 {
		return errors.New("bill not found")
	}
	return result.Error
}

// MarkAsPaid marks a bill as paid
func (s *BillService) MarkAsPaid(companyID, billID uuid.UUID, userID uuid.UUID, paidDate time.Time) (*models.Bill, error) {
	var bill models.Bill
	if err := s.db.Where("company_id = ? AND id = ?", companyID, billID).First(&bill).Error; err != nil {
		return nil, errors.New("bill not found")
	}

	if err := s.db.Model(&bill).Updates(map[string]interface{}{
		"status":    models.StatusPaid,
		"paid_date": paidDate,
	}).Error; err != nil {
		return nil, err
	}

	// Log activity
	s.logActivity(bill.ID, &userID, models.ActionStatusChanged, "Bill marked as paid")

	return s.GetByID(companyID, billID)
}

// GetActivities retrieves activity log for a bill
func (s *BillService) GetActivities(companyID, billID uuid.UUID) ([]models.BillActivity, error) {
	var activities []models.BillActivity
	err := s.db.
		Preload("User").
		Joins("JOIN bills ON bills.id = bill_activities.bill_id").
		Where("bills.company_id = ? AND bill_activities.bill_id = ?", companyID, billID).
		Order("bill_activities.created_at DESC").
		Find(&activities).Error

	return activities, err
}

// logActivity creates an activity log entry
func (s *BillService) logActivity(billID uuid.UUID, userID *uuid.UUID, action models.ActivityAction, details string) {
	activity := models.BillActivity{
		BillID:  billID,
		UserID:  userID,
		Action:  action,
		Details: &details,
	}
	s.db.Create(&activity)
}
