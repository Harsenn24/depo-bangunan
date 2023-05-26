package main

import (
	"depobangunan/app/config"
	"depobangunan/app/middleware"
	"depobangunan/app/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config.GetDBInstance()

	routers.RouterCustomer(router)

	router.Use(middleware.Authguard)

	routers.RouterOrder(router)
	
	router.Run(":4040")

}
