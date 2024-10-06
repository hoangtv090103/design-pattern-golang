package product

import "time"

type Product struct {
	ProductName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Factory Pattern here
func (p *Product) New() *Product { // Create an instance of that type
	product := Product{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &product
}
