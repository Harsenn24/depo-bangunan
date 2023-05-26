package intface

import "time"

type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Product    string    `json:"product"`
	Quantity   int       `json:"quantity"`
	Price      int       `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ListOrder struct {
	Id       int    `json:"id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Customer CustomerData
}

type CustomerData struct {
	Name string `json:"name"`
}

type UpadateOrder struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
