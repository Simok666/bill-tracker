package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Pagination holds pagination parameters
type Pagination struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Sort     string `json:"sort"`
	Order    string `json:"order"`
}

// GetPagination extracts pagination from query params
func GetPagination(c *gin.Context) Pagination {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	sort := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")
	if order != "asc" && order != "desc" {
		order = "desc"
	}

	return Pagination{
		Page:     page,
		PageSize: pageSize,
		Sort:     sort,
		Order:    order,
	}
}

// GetOffset calculates the database offset
func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

// GetOrderClause returns the GORM order clause
func (p *Pagination) GetOrderClause() string {
	return p.Sort + " " + p.Order
}
