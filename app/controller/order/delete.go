package order

import (
	"depobangunan/app/config"
	"depobangunan/app/responses"
	"strconv"
	"github.com/gin-gonic/gin"
)

func DeleteOrder(c *gin.Context) {
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

	result, err := db.Exec(`DELETE FROM orders WHERE id = $1`, id_int)
	
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
		c.JSON(400, responses.NewResponses(400, "Error", "Failed delete order"))
		return
	}

	c.JSON(200, responses.NewResponses(200, "Success", "Success delete order"))
}