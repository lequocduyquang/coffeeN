package main

type Discount struct {
	Name    string
	Percent int64
	Store   string
}

type IDiscount interface {
	setName(name string)
	setPercent(percent int64)
	setStore(store string)
	getName() string
	getPercent() int64
	getStore() string
}

func (d *Discount) setName(name string) {
	d.Name = name
}

func (d *Discount) setPercent(percent int64) {
	d.Percent = percent
}

func (d *Discount) setStore(store string) {
	d.Store = store
}

func (d *Discount) getName() string {
	return d.Name
}

func (d *Discount) getPercent() int64 {
	return d.Percent
}

func (d *Discount) getStore() string {
	return d.Store
}
