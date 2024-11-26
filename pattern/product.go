package pattern

import "time"

type Product struct {
	ProductNamge string
	CreateAt     time.Time
	UpdatedAt    time.Time
}

func (p *Product) New() *Product {
	product := Product{
		CreateAt:  time.Now(),
		UpdatedAt: time.Now(),
	}

	return &product
}
