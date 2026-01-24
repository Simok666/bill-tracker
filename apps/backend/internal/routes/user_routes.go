package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/dhani/bill-tracker-backend/internal/middleware"
	"github.com/dhani/bill-tracker-backend/internal/services"
	"github.com/dhani/bill-tracker-backend/internal/utils"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetProfile retrieves current user profile
// GET /api/users/profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	user := middleware.GetCurrentUser(c)

	profile, err := h.service.GetProfile(user.ID)
	if err != nil {
		utils.NotFound(c, "User not found")
		return
	}

	utils.Success(c, "", profile)
}

// UpdateProfile updates current user profile
// PUT /api/users/profile
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	user := middleware.GetCurrentUser(c)

	var input services.UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	profile, err := h.service.UpdateProfile(user.ID, input)
	if err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.Success(c, "Profile updated successfully", profile)
}

// ChangePassword changes current user password
// PUT /api/users/password
func (h *UserHandler) ChangePassword(c *gin.Context) {
	user := middleware.GetCurrentUser(c)

	var input services.ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	if err := h.service.ChangePassword(user.ID, input); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, "Password changed successfully", nil)
}
