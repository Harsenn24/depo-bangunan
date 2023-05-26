package customer

import (
	"depobangunan/app/config"
	"depobangunan/app/intface"
	"depobangunan/app/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateCustomer(c *gin.Context) {

	
	id := c.Param("id")

	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	var update_customer intface.UpdateCustomer

	if err := c.BindJSON(&update_customer); err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	if validationErr := validate.Struct(&update_customer); validationErr != nil {
		c.JSON(400, responses.NewResponses(400, "Error", validationErr.Error()))
		return
	}


	db, err := config.GetDBInstance()
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	result, err := db.Exec(`UPDATE customers SET name = $1, email = $2, updated_at = NOW() WHERE id = $3`, update_customer.Name, update_customer.Email, id_int)

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
		c.JSON(400, responses.NewResponses(400, "Error", "Failed update customer"))
		return
	}

	c.JSON(200, responses.NewResponses(200, "Success", "Success update customer"))

}
