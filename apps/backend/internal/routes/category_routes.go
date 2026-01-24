package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dhani/bill-tracker-backend/internal/middleware"
	"github.com/dhani/bill-tracker-backend/internal/services"
	"github.com/dhani/bill-tracker-backend/internal/utils"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// List retrieves all categories
// GET /api/categories
func (h *CategoryHandler) List(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	categories, err := h.service.List(companyID)
	if err != nil {
		utils.InternalError(c, "Failed to fetch categories")
		return
	}

	utils.Success(c, "", categories)
}

// GetByID retrieves a single category
// GET /api/categories/:id
func (h *CategoryHandler) GetByID(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	categoryID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "Invalid category ID")
		return
	}

	category, err := h.service.GetByID(companyID, categoryID)
	if err != nil {
		utils.NotFound(c, "Category not found")
		return
	}

	utils.Success(c, "", category)
}

// Create creates a new category
// POST /api/categories
func (h *CategoryHandler) Create(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	var input services.CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	category, err := h.service.Create(companyID, input)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.Created(c, "Category created successfully", category)
}

// Update updates an existing category
// PUT /api/categories/:id
func (h *CategoryHandler) Update(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	categoryID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "Invalid category ID")
		return
	}

	var input services.UpdateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	category, err := h.service.Update(companyID, categoryID, input)
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, "Category updated successfully", category)
}

// Delete deletes a category
// DELETE /api/categories/:id
func (h *CategoryHandler) Delete(c *gin.Context) {
	companyID := middleware.GetCompanyID(c)

	categoryID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "Invalid category ID")
		return
	}

	if err := h.service.Delete(companyID, categoryID); err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, "Category deleted successfully", nil)
}
