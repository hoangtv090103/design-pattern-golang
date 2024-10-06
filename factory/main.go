package main

import (
	"fmt"
	product "myapp/products"
)

func main() {
	factory := product.Product{}

	product := factory.New()

	fmt.Println("My product was created at", product.CreatedAt.UTC())
}
