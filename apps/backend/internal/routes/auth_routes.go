package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/dhani/bill-tracker-backend/internal/middleware"
	"github.com/dhani/bill-tracker-backend/internal/services"
	"github.com/dhani/bill-tracker-backend/internal/utils"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Register handles user registration
// POST /api/auth/register
func (h *AuthHandler) Register(c *gin.Context) {
	var input services.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	response, err := h.service.Register(input)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Created(c, "Registration successful", response)
}

// Login handles user login
// POST /api/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var input services.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	response, err := h.service.Login(input)
	if err != nil {
		utils.Unauthorized(c, err.Error())
		return
	}

	utils.Success(c, "Login successful", response)
}

// Logout handles user logout
// POST /api/auth/logout
func (h *AuthHandler) Logout(c *gin.Context) {
	// Get token from header
	token := c.GetHeader("Authorization")
	if len(token) > 7 {
		token = token[7:] // Remove "Bearer "
	}

	if err := h.service.Logout(token); err != nil {
		utils.InternalError(c, "Failed to logout")
		return
	}

	utils.Success(c, "Logout successful", nil)
}

// Me returns the current user's profile
// GET /api/auth/me
func (h *AuthHandler) Me(c *gin.Context) {
	user := middleware.GetCurrentUser(c)
	if user == nil {
		utils.Unauthorized(c, "Not authenticated")
		return
	}

	utils.Success(c, "", user)
}
