package order

import (
	"depobangunan/app/config"
	"depobangunan/app/intface"
	"depobangunan/app/responses"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListOrder(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	limit_page, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	offset := (page - 1) * limit_page

	db, err := config.GetDBInstance()
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	//query order join customer and output from customer only name

	result_query, err := db.Query("SELECT orders.id, customers.name, orders.price, orders.product, orders.quantity FROM orders INNER JOIN customers ON orders.customer_id = customers.id limit $1 offset $2", limit_page, offset)

	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	if result_query == nil {
		c.JSON(400, responses.NewResponses(400, "Error", "Data not found"))
		return
	}

	fmt.Println(result_query)

	defer result_query.Close()

	var list_orders []intface.ListOrder

	for result_query.Next() {
		var list_order intface.ListOrder
		var customerName string


		if err := result_query.Scan(&list_order.Id, &customerName, &list_order.Price, &list_order.Product, &list_order.Quantity); err != nil {
			c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
			return
		}

		list_order.Customer.Name = customerName

		list_orders = append(list_orders, list_order)
	}

	var count int
	error := db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&count)
	if error != nil {
		c.JSON(400, responses.NewResponses(400, "Error", error.Error()))
		return
	}

	c.JSON(200, responses.PaginationResponse(200, "Success" , page, limit_page, count, list_orders))



}



