package customer

import (
	"depobangunan/app/intface"

	"github.com/go-playground/validator/v10"
)

var Customer intface.Customer
var validate = validator.New()