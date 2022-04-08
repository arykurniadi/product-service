package healthcheck

import "dbo.id/product-service/models"

type IHealthCheckRepository interface {
	GetDBTimestamp() models.HealthCheck
}
