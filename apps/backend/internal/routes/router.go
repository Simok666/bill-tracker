package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/dhani/bill-tracker-backend/internal/middleware"
	"github.com/dhani/bill-tracker-backend/internal/services"
)

// SetupRouter configures all API routes
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.New()

	// Apply global middleware
	router.Use(gin.Recovery())
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Initialize services
	authService := services.NewAuthService(db)
	billService := services.NewBillService(db)
	vendorService := services.NewVendorService(db)
	categoryService := services.NewCategoryService(db)
	dashboardService := services.NewDashboardService(db)
	userService := services.NewUserService(db)

	// Initialize handlers
	authHandler := NewAuthHandler(authService)
	billHandler := NewBillHandler(billService)
	vendorHandler := NewVendorHandler(vendorService)
	categoryHandler := NewCategoryHandler(categoryService)
	dashboardHandler := NewDashboardHandler(dashboardService)
	userHandler := NewUserHandler(userService)

	// API routes
	api := router.Group("/api")
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", authHandler.Logout)
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Auth (protected)
			protected.GET("/auth/me", authHandler.Me)

			// Bills
			bills := protected.Group("/bills")
			{
				bills.GET("", billHandler.List)
				bills.GET("/:id", billHandler.GetByID)
				bills.POST("", billHandler.Create)
				bills.PUT("/:id", billHandler.Update)
				bills.DELETE("/:id", billHandler.Delete)
				bills.POST("/:id/pay", billHandler.Pay)
				bills.GET("/:id/activities", billHandler.GetActivities)
			}

			// Vendors
			vendors := protected.Group("/vendors")
			{
				vendors.GET("", vendorHandler.List)
				vendors.GET("/:id", vendorHandler.GetByID)
				vendors.POST("", vendorHandler.Create)
				vendors.PUT("/:id", vendorHandler.Update)
				vendors.DELETE("/:id", vendorHandler.Delete)
			}

			// Categories
			categories := protected.Group("/categories")
			{
				categories.GET("", categoryHandler.List)
				categories.GET("/:id", categoryHandler.GetByID)
				categories.POST("", categoryHandler.Create)
				categories.PUT("/:id", categoryHandler.Update)
				categories.DELETE("/:id", categoryHandler.Delete)
			}

			// Dashboard
			dashboard := protected.Group("/dashboard")
			{
				dashboard.GET("/stats", dashboardHandler.GetStats)
				dashboard.GET("/expenses-by-month", dashboardHandler.GetExpensesByMonth)
				dashboard.GET("/expenses-by-category", dashboardHandler.GetExpensesByCategory)
			}

			// Users
			users := protected.Group("/users")
			{
				users.GET("/profile", userHandler.GetProfile)
				users.PUT("/profile", userHandler.UpdateProfile)
				users.PUT("/password", userHandler.ChangePassword)
			}
		}
	}

	return router
}
