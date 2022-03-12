package main

import "fmt"

func main() {
	var coffeeBuilder = NewCoffeeBuilder()
	coffeeBuilder.SetName("Hot coffee")
	coffeeBuilder.SetImage("coffee.jpg")
	coffeeBuilder.SetSize("S")
	coffeeBuilder.SetPrice(30000)

	coffee, err := coffeeBuilder.Build()
	if err != nil {
		fmt.Println("Error creating coffee:", err)
	} else {
		fmt.Printf("Coffee: %+v\n", coffee)
	}
}
