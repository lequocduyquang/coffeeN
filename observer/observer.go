package main

import "fmt"

type Observer interface {
	onFinish(data string)
}

type DataListener struct {
	CustomerName string
}

func (dl *DataListener) onFinish(data string) {
	fmt.Println("Customer:", dl.CustomerName, " your order status:", data)
}
