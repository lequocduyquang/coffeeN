package main

import "fmt"

type DrinkType int

const (
	PhinCoffee DrinkType = iota
	Tea
	IceBlended
)

type Drink struct {
	Name  string
	Price float64
	Size  string
	Type  DrinkType
}

type Menu struct {
	Drinks []Drink
}

func (m *Menu) IterateDrinks(f func(Drink) error) {
	var err error
	for _, drink := range m.Drinks {
		err = f(drink)
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			break
		}
	}
}

func (m *Menu) CreateIterator() Iterator {
	return &DrinkIterator{
		drinks: m.Drinks,
	}
}
