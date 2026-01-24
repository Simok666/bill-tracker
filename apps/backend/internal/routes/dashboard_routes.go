package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/dhani/bill-tracker-backend/internal/middleware"
	"github.com/dhani/bill-tracker-backend/internal/services"
	"github.com/dhani/bill-tracker-backend/internal/utils"
)

type DashboardHandler struct {
	service *services.DashboardService
}

func NewDashboardHandler(service *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: service}
}

// GetStats retrieves KPI statistics
// GET /api/dashboard/stats
func (h *DashboardHandler) GetStats(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	stats, err := h.service.GetStats(companyID)
	if err != nil {
		utils.InternalError(c, "Failed to fetch statistics")
		return
	}

	utils.Success(c, "", stats)
}

// GetExpensesByMonth retrieves monthly expense breakdown
// GET /api/dashboard/expenses-by-month
func (h *DashboardHandler) GetExpensesByMonth(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	months, _ := strconv.Atoi(c.DefaultQuery("months", "12"))

	expenses, err := h.service.GetExpensesByMonth(companyID, months)
	if err != nil {
		utils.InternalError(c, "Failed to fetch monthly expenses")
		return
	}

	utils.Success(c, "", expenses)
}

// GetExpensesByCategory retrieves category expense breakdown
// GET /api/dashboard/expenses-by-category
func (h *DashboardHandler) GetExpensesByCategory(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	expenses, err := h.service.GetExpensesByCategory(companyID)
	if err != nil {
		utils.InternalError(c, "Failed to fetch category expenses")
		return
	}

	utils.Success(c, "", expenses)
}
