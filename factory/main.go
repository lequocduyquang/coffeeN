package main

import "fmt"

func main() {
	coupon1, _ := NewDiscount("coupon", "Coupon N", "CoffeeN", 20)
	voucher1, _ := NewDiscount("voucher", "Voucher N", "CoffeeN", 30)

	showInfo(coupon1)
	showInfo(voucher1)
}

func showInfo(discount IDiscount) {
	fmt.Printf("------  Show info ------\n")
	fmt.Printf("Name: %v\n", discount.getName())
	fmt.Printf("Store: %v\n", discount.getStore())
	fmt.Printf("Percent: %v\n", discount.getPercent())
}
