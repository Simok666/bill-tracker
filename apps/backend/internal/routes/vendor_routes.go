package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dhani/bill-tracker-backend/internal/middleware"
	"github.com/dhani/bill-tracker-backend/internal/services"
	"github.com/dhani/bill-tracker-backend/internal/utils"
)

type VendorHandler struct {
	service *services.VendorService
}

func NewVendorHandler(service *services.VendorService) *VendorHandler {
	return &VendorHandler{service: service}
}

// List retrieves all vendors
// GET /api/vendors
func (h *VendorHandler) List(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	vendors, err := h.service.List(companyID)
	if err != nil {
		utils.InternalError(c, "Failed to fetch vendors")
		return
	}

	utils.Success(c, "", vendors)
}

// GetByID retrieves a single vendor
// GET /api/vendors/:id
func (h *VendorHandler) GetByID(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	vendorID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "Invalid vendor ID")
		return
	}

	vendor, err := h.service.GetByID(companyID, vendorID)
	if err != nil {
		utils.NotFound(c, "Vendor not found")
		return
	}

	utils.Success(c, "", vendor)
}

// Create creates a new vendor
// POST /api/vendors
func (h *VendorHandler) Create(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	var input services.CreateVendorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	vendor, err := h.service.Create(companyID, input)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.Created(c, "Vendor created successfully", vendor)
}

// Update updates an existing vendor
// PUT /api/vendors/:id
func (h *VendorHandler) Update(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	vendorID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "Invalid vendor ID")
		return
	}

	var input services.UpdateVendorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	vendor, err := h.service.Update(companyID, vendorID, input)
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, "Vendor updated successfully", vendor)
}

// Delete deletes a vendor
// DELETE /api/vendors/:id
func (h *VendorHandler) Delete(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	vendorID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "Invalid vendor ID")
		return
	}

	if err := h.service.Delete(companyID, vendorID); err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, "Vendor deleted successfully", nil)
}
