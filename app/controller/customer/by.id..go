package customer

import (
	"depobangunan/app/config"
	"depobangunan/app/intface"
	"depobangunan/app/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DetailCosutomer(c *gin.Context) {
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

	find_costumer := db.QueryRow(`SELECT id, name, email FROM customers WHERE id = $1`, id_int)

	var customer_by_id intface.ListCustomer
	
	if err := find_costumer.Scan(&customer_by_id.Id, &customer_by_id.Name, &customer_by_id.Email); err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	c.JSON(200, responses.NewResponses(200, "Success", customer_by_id))
}
