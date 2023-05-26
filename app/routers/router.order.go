package routers

import (
	"depobangunan/app/controller/order"
	"github.com/gin-gonic/gin"
)

func RouterOrder(router *gin.Engine) {

	router.POST("/order", order.CreateOrder)

	router.GET("/order", order.ListOrder)

	router.GET("/order/:id", order.OrderById)

	router.PUT("/order/:id", order.UpdatOrder)

	router.DELETE("/order/:id", order.DeleteOrder)

}
