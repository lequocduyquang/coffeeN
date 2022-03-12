package main

import "fmt"

type CoffeeMachine struct {
	BeanAmount  float64
	MilkAmount  float64
	WaterAmount float64
	WaterTemp   int
}

func (c *CoffeeMachine) StartCoffee(beanAmount float64) {
	c.BeanAmount = beanAmount
	fmt.Printf("Starting coffee with beans: %v\n", beanAmount)
}

func (c *CoffeeMachine) GrindBeans() bool {
	fmt.Printf("Grinding beans: %v\n", c.BeanAmount)
	return true
}

func (c *CoffeeMachine) SetWaterTemp(temp int) {
	c.WaterTemp = temp
	fmt.Printf("Set water temp: %v\n", temp)
}

func (c *CoffeeMachine) UseHotWater(amount float64) float64 {
	c.WaterAmount = amount
	fmt.Printf("Add hot water: %v\n", amount)
	return amount
}

func (c *CoffeeMachine) UseMilk(amount float64) float64 {
	fmt.Printf("Add milk: %v\n", amount)
	c.MilkAmount = amount
	return amount
}

func (c *CoffeeMachine) EndCoffee() {
	fmt.Println("Finish!!!!")
}
