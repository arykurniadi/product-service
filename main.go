package main

import (
	"fmt"
	"log"

	"dbo.id/product-service/config"
	gorm "dbo.id/product-service/db"
	"github.com/gin-gonic/gin"

	routes "dbo.id/product-service/app"
	CustomerRepository "dbo.id/product-service/app/customer/repository"
	HCRepository "dbo.id/product-service/app/health-check/repository"
	OrderRepository "dbo.id/product-service/app/order/repository"
	UserRepository "dbo.id/product-service/app/user/repository"

	CustomerUsecase "dbo.id/product-service/app/customer/usecase"
	HCUsecase "dbo.id/product-service/app/health-check/usecase"
	OrderUsecase "dbo.id/product-service/app/order/usecase"
	UserUsecase "dbo.id/product-service/app/user/usecase"
)

var appConfig = config.Config.App

func main() {
	r := gin.New()

	db := gorm.MysqlConn()

	r.Use(gin.Recovery())

	hcr := HCRepository.NewHealthCheckRepository(db)
	userRepo := UserRepository.NewUserRepository(db)
	customerRepo := CustomerRepository.NewCustomerRepository(db)
	orderRepo := OrderRepository.NewOrderRepository(db)

	hcu := HCUsecase.NewHealthCheckUsecase(hcr)
	user := UserUsecase.NewUserUsecase(userRepo)
	customer := CustomerUsecase.NewCustomerUsecase(customerRepo)
	order := OrderUsecase.NewOrderUsecase(orderRepo, customerRepo)

	routes.HealthCheckHttpHandler(r, hcu)
	routes.UserHttpHandler(r, user)
	routes.CustomerHttpHandler(r, customer)
	routes.OrderHttpHandler(r, order)

	// check auth endpoint for admin access
	// r.Use(AuthAdminMiddleware())

	if err := r.Run(fmt.Sprintf(":%s", appConfig.HTTPPort)); err != nil {
		log.Fatal(err)
	}
}

func AuthAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, hasAuth := c.Request.BasicAuth()
		if !hasAuth == true && username != "admin" && password != "admin" {
			c.AbortWithStatus(401)
		}

		c.Next()
	}
}
