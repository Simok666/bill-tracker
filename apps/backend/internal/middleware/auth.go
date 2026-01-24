package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/dhani/bill-tracker-backend/internal/config"
	"github.com/dhani/bill-tracker-backend/internal/database"
	"github.com/dhani/bill-tracker-backend/internal/models"
	"github.com/dhani/bill-tracker-backend/internal/utils"
)

// Claims represents JWT claims
type Claims struct {
	UserID    uuid.UUID `json:"user_id"`
	CompanyID uuid.UUID `json:"company_id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(time.Duration(config.AppConfig.JWTExpirationHours) * time.Hour)

	claims := &Claims{
		UserID:    user.ID,
		CompanyID: user.CompanyID,
		Email:     user.Email,
		Role:      string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.ID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

// ValidateToken validates and parses a JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

// AuthMiddleware validates JWT tokens and sets user context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Unauthorized(c, "Authorization header required")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Unauthorized(c, "Invalid authorization header format")
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := ValidateToken(tokenString)
		if err != nil {
			utils.Unauthorized(c, "Invalid or expired token")
			c.Abort()
			return
		}

		// Load user from database
		var user models.User
		if err := database.DB.First(&user, "id = ?", claims.UserID).Error; err != nil {
			utils.Unauthorized(c, "User not found")
			c.Abort()
			return
		}

		// Set user in context
		c.Set("user", &user)
		c.Set("user_id", user.ID)
		c.Set("company_id", user.CompanyID)

		c.Next()
	}
}

// GetCurrentUser retrieves the authenticated user from context
func GetCurrentUser(c *gin.Context) *models.User {
	user, exists := c.Get("user")
	if !exists {
		return nil
	}
	return user.(*models.User)
}

// GetCompanyID retrieves the company ID from context
func GetCompanyID(c *gin.Context) uuid.UUID {
	companyID, exists := c.Get("company_id")
	if !exists {
		return uuid.Nil
	}
	return companyID.(uuid.UUID)
}

// AdminOnly middleware restricts access to admin users
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := GetCurrentUser(c)
		if user == nil || user.Role != models.RoleAdmin {
			utils.Forbidden(c, "Admin access required")
			c.Abort()
			return
		}
		c.Next()
	}
}
