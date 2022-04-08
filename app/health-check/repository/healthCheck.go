package repository

import (
	HealthCheckInterface "dbo.id/product-service/app/health-check"
	"dbo.id/product-service/models"
	"github.com/jinzhu/gorm"
)

// Define your repository connection
// In this case, we use a SQL connection and use gorm as an ORM
type HealthCheckRepository struct {
	Conn *gorm.DB
}

// Define an exported function. Call this in main.go and bind your connection to the struct
func NewHealthCheckRepository(Conn *gorm.DB) HealthCheckInterface.IHealthCheckRepository {
	return &HealthCheckRepository{Conn}
}

// Define your custom functions
// Please don't put any logic on the repository
// Functions on the repository just help you to the CRUD operation
// To export the functions, you need to register function to the interface

func (m *HealthCheckRepository) GetDBTimestamp() models.HealthCheck {
	var healthCheck models.HealthCheck

	tx := m.Conn.Begin()
	tx.Raw("SELECT current_timestamp").Scan(&healthCheck)
	tx.Commit()

	return healthCheck
}
