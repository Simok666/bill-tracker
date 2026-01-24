package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dhani/bill-tracker-backend/internal/middleware"
	"github.com/dhani/bill-tracker-backend/internal/services"
	"github.com/dhani/bill-tracker-backend/internal/utils"
)

type BillHandler struct {
	service *services.BillService
}

func NewBillHandler(service *services.BillService) *BillHandler {
	return &BillHandler{service: service}
}

// List retrieves bills with pagination and filters
// GET /api/bills
func (h *BillHandler) List(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)
	
	filters := services.BillFilters{
		Status:     c.Query("status"),
		Search:     c.Query("search"),
		Pagination: utils.GetPagination(c),
	}

	bills, total, err := h.service.List(companyID, filters)
	if err != nil {
		utils.InternalError(c, "Failed to fetch bills")
		return
	}

	utils.Paginated(c, bills, total, filters.Pagination.Page, filters.Pagination.PageSize)
}

// GetByID retrieves a single bill
// GET /api/bills/:id
func (h *BillHandler) GetByID(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)
	
	billID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "Invalid bill ID")
		return
	}

	bill, err := h.service.GetByID(companyID, billID)
	if err != nil {
		utils.NotFound(c, "Bill not found")
		return
	}

	utils.Success(c, "", bill)
}

// Create creates a new bill
// POST /api/bills
func (h *BillHandler) Create(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)
	user := middleware.GetCurrentUser(c)

	var input services.CreateBillInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	bill, err := h.service.Create(companyID, user.ID, input)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.Created(c, "Bill created successfully", bill)
}

// Update updates an existing bill
// PUT /api/bills/:id
func (h *BillHandler) Update(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)
	user := middleware.GetCurrentUser(c)

	billID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "Invalid bill ID")
		return
	}

	var input services.UpdateBillInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	bill, err := h.service.Update(companyID, billID, user.ID, input)
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, "Bill updated successfully", bill)
}

// Delete deletes a bill
// DELETE /api/bills/:id
func (h *BillHandler) Delete(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	billID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "Invalid bill ID")
		return
	}

	if err := h.service.Delete(companyID, billID); err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, "Bill deleted successfully", nil)
}

// Pay marks a bill as paid
// POST /api/bills/:id/pay
func (h *BillHandler) Pay(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)
	user := middleware.GetCurrentUser(c)

	billID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "Invalid bill ID")
		return
	}

	var input struct {
		PaidDate *time.Time `json:"paid_date"`
	}
	c.ShouldBindJSON(&input)

	paidDate := time.Now()
	if input.PaidDate != nil {
		paidDate = *input.PaidDate
	}

	bill, err := h.service.MarkAsPaid(companyID, billID, user.ID, paidDate)
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, "Bill marked as paid", bill)
}

// GetActivities retrieves activity log for a bill
// GET /api/bills/:id/activities
func (h *BillHandler) GetActivities(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	billID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "Invalid bill ID")
		return
	}

	activities, err := h.service.GetActivities(companyID, billID)
	if err != nil {
		utils.InternalError(c, "Failed to fetch activities")
		return
	}

	utils.Success(c, "", activities)
}
