package order

import (
	"depobangunan/app/config"
	"depobangunan/app/intface"
	"depobangunan/app/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdatOrder(c *gin.Context) {
	id := c.Param("id")

	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	var update_order intface.UpadateOrder

	if err := c.BindJSON(&update_order); err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	if validationErr := validate.Struct(&update_order); validationErr != nil {
		c.JSON(400, responses.NewResponses(400, "Error", validationErr.Error()))
		return
	}

	db, err := config.GetDBInstance()
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	result, err := db.Exec(`UPDATE orders SET product = $1, quantity = $2, price = $3, updated_at = NOW() WHERE id = $4`, update_order.Product, update_order.Quantity, update_order.Price, id_int)

	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	row_affected, err := result.RowsAffected()

	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	if row_affected == 0 {
		c.JSON(400, responses.NewResponses(400, "Error", "Failed update order"))
		return
	}

	c.JSON(200, responses.NewResponses(200, "Success", "Success update order"))

}