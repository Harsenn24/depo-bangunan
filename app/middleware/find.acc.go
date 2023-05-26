package middleware

import (
	"depobangunan/app/config"
	"depobangunan/app/intface"
	"depobangunan/app/responses"

	"github.com/gin-gonic/gin"
)

func FindAccount(c *gin.Context, email string) ([]intface.ListCustomer, error) {
	db, err := config.GetDBInstance()
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return nil, err
	}

	find_account, err := db.Query(`SELECT id, name, email FROM customers WHERE email = $1`, email)
	
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return nil, err
	}
	
	defer find_account.Close()

	var CheckAccount []intface.ListCustomer

	for find_account.Next() {
		var account intface.ListCustomer

		if err := find_account.Scan(&account.Id, &account.Name, &account.Email); err != nil {
			c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
			return nil, err
		}

		CheckAccount = append(CheckAccount, account)
	}

	return CheckAccount, nil
}
