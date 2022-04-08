package healthcheck

import "dbo.id/product-service/models"

type IHealthCheckUsecase interface {
	GetDBTimestamp() models.HealthCheck
}
