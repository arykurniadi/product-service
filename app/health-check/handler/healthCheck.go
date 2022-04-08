package handler

import (
	"time"

	Base "dbo.id/product-service/app/api/handler"
	HealthCheckInterface "dbo.id/product-service/app/health-check"
	"github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
	HealthStatus string    `json:"health_status"`
	DBTimestamp  time.Time `json:"database_timestamp"`
}

type HealthCheckHandler struct {
	HealthCheckUsecase HealthCheckInterface.IHealthCheckUsecase
}

// Handler just handle how data come and how data will serve
// To process the incoming data, you need to connect to the usecase via interface

func (a *HealthCheckHandler) Check(c *gin.Context) {
	healthCheck := a.HealthCheckUsecase.GetDBTimestamp()
	res := &HealthCheckResponse{
		HealthStatus: "GOOD",
		DBTimestamp:  healthCheck.CurrentTimestamp,
	}

	Base.RespondJSON(c, res, nil)
	return
}
