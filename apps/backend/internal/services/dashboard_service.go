package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/dhani/bill-tracker-backend/internal/models"
)

type DashboardService struct {
	db *gorm.DB
}

func NewDashboardService(db *gorm.DB) *DashboardService {
	return &DashboardService{db: db}
}

// DashboardStats holds KPI data
type DashboardStats struct {
	TotalExpense         decimal.Decimal `json:"total_expense"`
	PaidThisMonth        decimal.Decimal `json:"paid_this_month"`
	UnpaidAmount         decimal.Decimal `json:"unpaid_amount"`
	OverdueBillsCount    int64           `json:"overdue_bills_count"`
	ExpenseChangePercent float64         `json:"expense_change_percent"`
}

// MonthlyExpense holds monthly expense data
type MonthlyExpense struct {
	Month  string          `json:"month"`
	Amount decimal.Decimal `json:"amount"`
}

// CategoryExpense holds category expense data
type CategoryExpense struct {
	CategoryID   uuid.UUID       `json:"category_id"`
	CategoryName string          `json:"category_name"`
	Amount       decimal.Decimal `json:"amount"`
	Percentage   float64         `json:"percentage"`
}

// GetStats retrieves dashboard KPI statistics
func (s *DashboardService) GetStats(companyID uuid.UUID) (*DashboardStats, error) {
	stats := &DashboardStats{}

	// Total expense (all paid bills)
	var totalExpense struct {
		Total decimal.Decimal
	}
	s.db.Model(&models.Bill{}).
		Select("COALESCE(SUM(amount), 0) as total").
		Where("company_id = ? AND status = ?", companyID, models.StatusPaid).
		Scan(&totalExpense)
	stats.TotalExpense = totalExpense.Total

	// Paid this month
	startOfMonth := time.Now().UTC().Truncate(24*time.Hour).AddDate(0, 0, -time.Now().Day()+1)
	var paidThisMonth struct {
		Total decimal.Decimal
	}
	s.db.Model(&models.Bill{}).
		Select("COALESCE(SUM(amount), 0) as total").
		Where("company_id = ? AND status = ? AND paid_date >= ?", companyID, models.StatusPaid, startOfMonth).
		Scan(&paidThisMonth)
	stats.PaidThisMonth = paidThisMonth.Total

	// Unpaid amount
	var unpaidAmount struct {
		Total decimal.Decimal
	}
	s.db.Model(&models.Bill{}).
		Select("COALESCE(SUM(amount), 0) as total").
		Where("company_id = ? AND status IN (?, ?)", companyID, models.StatusUnpaid, models.StatusOverdue).
		Scan(&unpaidAmount)
	stats.UnpaidAmount = unpaidAmount.Total

	// Overdue bills count
	s.db.Model(&models.Bill{}).
		Where("company_id = ? AND status = ?", companyID, models.StatusOverdue).
		Count(&stats.OverdueBillsCount)

	// Also count bills that are past due but not marked overdue
	var additionalOverdue int64
	s.db.Model(&models.Bill{}).
		Where("company_id = ? AND status = ? AND due_date < ?", companyID, models.StatusUnpaid, time.Now()).
		Count(&additionalOverdue)
	stats.OverdueBillsCount += additionalOverdue

	// Expense change percent (compare last month to this month)
	startOfLastMonth := startOfMonth.AddDate(0, -1, 0)
	var lastMonthTotal struct {
		Total decimal.Decimal
	}
	s.db.Model(&models.Bill{}).
		Select("COALESCE(SUM(amount), 0) as total").
		Where("company_id = ? AND status = ? AND paid_date >= ? AND paid_date < ?",
			companyID, models.StatusPaid, startOfLastMonth, startOfMonth).
		Scan(&lastMonthTotal)

	if !lastMonthTotal.Total.IsZero() {
		change := stats.PaidThisMonth.Sub(lastMonthTotal.Total)
		changePercent, _ := change.Div(lastMonthTotal.Total).Mul(decimal.NewFromInt(100)).Float64()
		stats.ExpenseChangePercent = changePercent
	}

	return stats, nil
}

// GetExpensesByMonth retrieves monthly expense breakdown
func (s *DashboardService) GetExpensesByMonth(companyID uuid.UUID, months int) ([]MonthlyExpense, error) {
	if months <= 0 {
		months = 12
	}

	var results []MonthlyExpense

	startDate := time.Now().AddDate(0, -months+1, 0).Truncate(24 * time.Hour)
	startDate = startDate.AddDate(0, 0, -startDate.Day()+1)

	rows, err := s.db.Model(&models.Bill{}).
		Select("TO_CHAR(paid_date, 'YYYY-MM') as month, COALESCE(SUM(amount), 0) as amount").
		Where("company_id = ? AND status = ? AND paid_date >= ?", companyID, models.StatusPaid, startDate).
		Group("TO_CHAR(paid_date, 'YYYY-MM')").
		Order("month ASC").
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var expense MonthlyExpense
		if err := rows.Scan(&expense.Month, &expense.Amount); err != nil {
			return nil, err
		}
		results = append(results, expense)
	}

	return results, nil
}

// GetExpensesByCategory retrieves category expense breakdown
func (s *DashboardService) GetExpensesByCategory(companyID uuid.UUID) ([]CategoryExpense, error) {
	var results []CategoryExpense

	rows, err := s.db.Model(&models.Bill{}).
		Select("bills.category_id, categories.name as category_name, COALESCE(SUM(bills.amount), 0) as amount").
		Joins("LEFT JOIN categories ON categories.id = bills.category_id").
		Where("bills.company_id = ? AND bills.status = ?", companyID, models.StatusPaid).
		Group("bills.category_id, categories.name").
		Order("amount DESC").
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var total decimal.Decimal
	var rawResults []CategoryExpense

	for rows.Next() {
		var expense CategoryExpense
		if err := rows.Scan(&expense.CategoryID, &expense.CategoryName, &expense.Amount); err != nil {
			return nil, err
		}
		rawResults = append(rawResults, expense)
		total = total.Add(expense.Amount)
	}

	// Calculate percentages
	for _, r := range rawResults {
		pct := float64(0)
		if !total.IsZero() {
			pctDec, _ := r.Amount.Div(total).Mul(decimal.NewFromInt(100)).Float64()
			pct = pctDec
		}
		results = append(results, CategoryExpense{
			CategoryID:   r.CategoryID,
			CategoryName: r.CategoryName,
			Amount:       r.Amount,
			Percentage:   pct,
		})
	}

	return results, nil
}
