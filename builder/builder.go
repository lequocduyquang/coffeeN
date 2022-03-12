package main

import "fmt"

type CoffeeBuilder struct {
	Name  string
	Image string
	Size  string
	Price float64
}

func NewCoffeeBuilder() *CoffeeBuilder {
	return &CoffeeBuilder{}
}

func (cb *CoffeeBuilder) SetName(name string) {
	cb.Name = name
}

func (cb *CoffeeBuilder) SetImage(image string) {
	cb.Image = image
}

func (cb *CoffeeBuilder) SetSize(size string) {
	cb.Size = size
}

func (cb *CoffeeBuilder) SetPrice(price float64) {
	cb.Price = price
}

func (cb *CoffeeBuilder) Build() (*Coffee, error) {
	if cb.Size == "S" && cb.Price != 30000 {
		return nil, fmt.Errorf("Price of size S must be 30.000")
	}
	if cb.Size == "M" && cb.Price != 40000 {
		return nil, fmt.Errorf("rice of size M must be 40.000")
	}
	if cb.Size == "L" && cb.Price != 50000 {
		return nil, fmt.Errorf("rice of size L must be 50.000")
	}

	return &Coffee{
		Name:  cb.Name,
		Size:  cb.Size,
		Image: cb.Image,
		Price: cb.Price,
	}, nil
}
