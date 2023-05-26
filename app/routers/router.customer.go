package routers

import (
	"depobangunan/app/controller/customer"
	"github.com/gin-gonic/gin"
)

func RouterCustomer(router *gin.Engine) {

	router.POST("/customer", customer.CreateCustomer)

	router.POST("/login", customer.LoginCustomer)

	router.GET("/customer", customer.ListCustomer)

	router.GET("/customer/:id", customer.DetailCosutomer)

	router.PUT("/customer/:id", customer.UpdateCustomer)

	router.DELETE("/customer/:id", customer.DeleteCustomer)

}
