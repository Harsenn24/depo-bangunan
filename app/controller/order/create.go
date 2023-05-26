package order

import (
	"depobangunan/app/config"
	"depobangunan/app/intface"
	"depobangunan/app/responses"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	payload_jwt := c.MustGet("user").(*intface.JwtClaim)

	db, err := config.GetDBInstance()
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	if err := c.BindJSON(&Order); err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	if validationErr := validate.Struct(&Order); validationErr != nil {
		c.JSON(400, responses.NewResponses(400, "Error", validationErr.Error()))
		return
	}

	insert_oder := db.QueryRow("INSERT INTO orders (customer_id, product, quantity, price) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at", payload_jwt.Id, Order.Product, Order.Quantity, Order.Price)

	if err := insert_oder.Scan(&Order.ID, &Order.CreatedAt, &Order.UpdatedAt); err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	c.JSON(201, responses.NewResponses(201, "Success", "Success create order"))
}
