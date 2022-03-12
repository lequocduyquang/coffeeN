package main

import "fmt"

func MakeHotCoffee(size float64) {
	fmt.Println("Hot coffee in process ...")

	hotCoffee := &CoffeeMachine{}

	beansAmount := (size / 8.0) * 5.0
	hotCoffee.StartCoffee(beansAmount)
	hotCoffee.GrindBeans()
	hotCoffee.SetWaterTemp(100)
	hotCoffee.UseHotWater(size)
	hotCoffee.EndCoffee()
}

func MakeMilkCoffee(size float64) {
	fmt.Println("Milk coffee in process ...")

	milkCoffee := &CoffeeMachine{}

	beansAmount := (size / 8.0) * 5.0
	milkCoffee.StartCoffee(beansAmount)
	milkCoffee.GrindBeans()
	milkCoffee.SetWaterTemp(100)
	milkCoffee.UseHotWater(size)

	milk := (size / 8.0) * 2.0
	milkCoffee.UseMilk(milk)
	milkCoffee.EndCoffee()
}
