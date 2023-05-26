package order

import (
	"depobangunan/app/config"
	"depobangunan/app/intface"
	"depobangunan/app/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OrderById(c *gin.Context) {
	id := c.Param("id")

	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	db, err := config.GetDBInstance()
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	find_order := db.QueryRow(`SELECT orders.id, customers.name, orders.quantity, orders.price, orders.product FROM orders JOIN customers ON orders.customer_id = customers.id WHERE orders.id = $1`, id_int)

	var order_by_id intface.ListOrder

	if err := find_order.Scan(&order_by_id.Id, &order_by_id.Customer.Name, &order_by_id.Quantity, &order_by_id.Price, &order_by_id.Product); err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	c.JSON(200, responses.NewResponses(200, "Success", order_by_id))

	
}