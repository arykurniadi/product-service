package app

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	CustomerInterface "dbo.id/product-service/app/customer"
	HealthCheckInterface "dbo.id/product-service/app/health-check"
	OrderInterface "dbo.id/product-service/app/order"
	UserInterface "dbo.id/product-service/app/user"
	"dbo.id/product-service/libraries"

	CustomerHandler "dbo.id/product-service/app/customer/handler"
	HCHandler "dbo.id/product-service/app/health-check/handler"
	OrderHandler "dbo.id/product-service/app/order/handler"
	UserHandler "dbo.id/product-service/app/user/handler"
)

func HealthCheckHttpHandler(r *gin.Engine, us HealthCheckInterface.IHealthCheckUsecase) {
	handler := &HCHandler.HealthCheckHandler{
		HealthCheckUsecase: us,
	}

	route := r.Group("/test")
	route.GET("/health-check", handler.Check)
}

func UserHttpHandler(r *gin.Engine, us UserInterface.IUserUsecase) {
	handler := &UserHandler.UserHandler{
		UserUsecase: us,
	}

	route := r.Group("/user")
	route.GET("/list", handler.GetListUser)
	route.GET("/detail/:id", handler.GetUserById)
	route.POST("", handler.Create)
	route.PUT("/:id", handler.Update)
	route.DELETE("/:id", handler.Delete)
}

func CustomerHttpHandler(r *gin.Engine, cs CustomerInterface.ICustomerUsecase) {
	handler := &CustomerHandler.CustomerHandler{
		CustomerUsercase: cs,
	}

	route := r.Group("/customer").Use(JwtMiddleware())
	route.GET("/list", handler.GetListCustomer)
	route.GET("/detail/:id", handler.GetCustomerById)
	route.POST("", handler.Create)
	route.PUT("/:id", handler.Update)
	route.DELETE("/:id", handler.Delete)
}

func OrderHttpHandler(r *gin.Engine, od OrderInterface.IOrderUsecase) {
	handler := &OrderHandler.OrderHandler{
		OrderUsecase: od,
	}

	route := r.Group("/order").Use(JwtMiddleware())
	route.GET("/list", handler.GetListOrder)
	route.GET("/detail/:id", handler.GetOrderById)
	route.POST("", handler.Create)
}

func JwtMiddleware() gin.HandlerFunc {
	errorMessage := struct {
		Message string `json:"message"`
	}{
		Message: "Unauthorized",
	}

	return func(c *gin.Context) {
		var jwtService libraries.JWTService = libraries.JWTAuthService()
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.Request.Header.Get("authorization")
		if govalidator.IsNull(authHeader) || !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(401, errorMessage)
			return
		}

		tokenString := strings.Split(authHeader, " ")
		token, err := jwtService.ValidateToken(tokenString[1])
		if err != nil {
			c.AbortWithStatusJSON(401, errorMessage)
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			if claims["user"] != true {
				c.AbortWithStatusJSON(401, errorMessage)
				return
			}
		} else {
			c.AbortWithStatusJSON(401, errorMessage)
			return
		}
		c.Next()
	}
}
