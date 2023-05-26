package customer

import (
	"depobangunan/app/config"
	"depobangunan/app/helper"
	"depobangunan/app/responses"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	db, err := config.GetDBInstance()
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	if err := c.BindJSON(&Customer); err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	if validationErr := validate.Struct(&Customer); validationErr != nil {
		c.JSON(400, responses.NewResponses(400, "Error", validationErr.Error()))
		return
	}

	hash_password, err := helper.HashPassword(Customer.Password)
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	insert_customer := db.QueryRow("INSERT INTO customers (name, password, email) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at", Customer.Name, hash_password, Customer.Email)

	if err := insert_customer.Scan(&Customer.ID, &Customer.CreatedAt, &Customer.UpdatedAt); err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	c.JSON(201, responses.NewResponses(201, "Success", "Success create customer"))
}
