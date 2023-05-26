package middleware

import (
	"depobangunan/app/environment"
	"depobangunan/app/helper"
	"depobangunan/app/responses"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Authguard(c *gin.Context) {

	author := c.GetHeader("authorization")

	environment.ExportEnv()
	keyJwt := os.Getenv("KEYJWT")

	secretKey := []byte(keyJwt)

	data, err := helper.DecryptJWT(author, secretKey)

	if err != nil {
		c.AbortWithStatusJSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	email := data.Email

	find_account, err := FindAccount(c, email)

	if err != nil {
		c.AbortWithStatusJSON(400, responses.NewResponses(400, "Error", err.Error()))
		return
	}

	fmt.Println(len(find_account))

	if len(find_account) != 1 {
		c.AbortWithStatusJSON(400, responses.NewResponses(400, "Error", "Email not found"))
		return
	} else {

		c.Set("user", data)

		c.Next()
	}

}
