package customer

import (
	"depobangunan/app/config"
	"depobangunan/app/helper"
	"depobangunan/app/intface"
	"depobangunan/app/responses"

	"github.com/gin-gonic/gin"
)

func LoginCustomer(c *gin.Context) {

	var userlogin intface.LoginCustomer

	var CheckAccount intface.CheckAccount

	db, err := config.GetDBInstance()
	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	if err := c.BindJSON(&userlogin); err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	if validationErr := validate.Struct(&userlogin); validationErr != nil {
		c.JSON(400, responses.NewResponses(400, "Error", validationErr.Error()))
		return
	}

	find_user := db.QueryRow(`SELECT id, email, password FROM customers WHERE email = $1`, userlogin.Email)

	if err := find_user.Scan(&CheckAccount.Id, &CheckAccount.Email, &CheckAccount.Password); err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", "Email not found"))
		return
	}

	matchPassword := helper.DecryptPassword(CheckAccount.Password, userlogin.Password)

	if !matchPassword {
		c.JSON(400, responses.NewResponses(400, "Error", "Password not match"))
		return
	}

	token, err := helper.JwtSign(&CheckAccount)

	if err != nil {
		c.JSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	c.JSON(200, responses.NewResponses(200, "Success", token))

}
