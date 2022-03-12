package main

import "fmt"

var menu *Menu = &Menu{
	Drinks: []Drink{
		{
			Name:  "Phin sữa đá",
			Price: 30000,
			Size:  "10oz",
			Type:  PhinCoffee,
		},
		{
			Name:  "Trà sen vàng",
			Price: 39000,
			Size:  "10oz",
			Type:  Tea,
		},
		{
			Name:  "Freeze trà xanh",
			Price: 49000,
			Size:  "10oz",
			Type:  IceBlended,
		},
		{
			Name:  "Bạc xỉu đá",
			Price: 35000,
			Size:  "12oz",
			Type:  PhinCoffee,
		},
		{
			Name:  "Trà thạch đào",
			Price: 49000,
			Size:  "12oz",
			Type:  Tea,
		},
		{
			Name:  "Freeze Sô-cô-la",
			Price: 59000,
			Size:  "12oz",
			Type:  IceBlended,
		},
	},
}

func main() {
	menu.IterateDrinks(func(d Drink) error {
		fmt.Printf("Drink name: %v\n", d.Name)
		return nil
	})

	iterator := menu.CreateIterator()
	for iterator.hasNext() {
		drink := iterator.next()
		fmt.Printf("Drink: %+v", drink)
	}
}
