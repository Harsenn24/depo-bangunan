package customer

import (
	"depobangunan/app/config"
	"depobangunan/app/intface"
	"depobangunan/app/responses"
	"strconv"
	"github.com/gin-gonic/gin"
)

func ListCustomer(c *gin.Context) {

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

	result_query, err := db.Query("SELECT id, name, email FROM customers limit $1 offset $2", limit_page, offset)
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	if result_query == nil {
		c.JSON(400, responses.NewResponses(400, "Error", "Data not found"))
		return
	}

	defer result_query.Close()

	var list_customers []intface.ListCustomer

	for result_query.Next() {
		var customer intface.ListCustomer

		if err := result_query.Scan(&customer.Id, &customer.Name, &customer.Email); err != nil {
			c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
			return
		}

		list_customers = append(list_customers, customer)
	}

	var count int
	error := db.QueryRow("SELECT COUNT(*) FROM customers").Scan(&count)
	if error != nil {
		c.JSON(400, responses.NewResponses(400, "Error", error.Error()))	
		return
	}

	c.JSON(200, responses.PaginationResponse(200, "Success", page, limit_page, count, list_customers))

}
