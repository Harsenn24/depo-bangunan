package intface

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CheckAccount struct {
	Email    string
	Id       string
	Password string
}

type JwtClaim struct {
	Email    string
	Password string
	Id       string
	jwt.StandardClaims
}

type LoginCustomer struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ListCustomer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateCustomer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
