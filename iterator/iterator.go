package main

type Iterator interface {
	hasNext() bool
	next() *Drink
}

type IterableDrinks interface {
	createIterator() Iterator
}

type DrinkIterator struct {
	current int
	drinks  []Drink
}

func (d *DrinkIterator) hasNext() bool {
	if d.current < len(d.drinks) {
		return true
	}
	return false
}

func (d *DrinkIterator) next() *Drink {
	if d.hasNext() {
		drink := d.drinks[d.current]
		d.current++
		return &drink
	}
	return nil
}
