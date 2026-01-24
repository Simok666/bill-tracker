package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/dhani/bill-tracker-backend/internal/config"
)

// CORSMiddleware configures CORS for the API
func CORSMiddleware() gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowOrigins:     []string{config.AppConfig.FrontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 hours
	}

	return cors.New(corsConfig)
}
