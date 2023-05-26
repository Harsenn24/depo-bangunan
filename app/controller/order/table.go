package order

import (
	"depobangunan/app/intface"

	"github.com/go-playground/validator/v10"
)

var Order intface.Order
var validate = validator.New()