package main

type Coupon struct {
	Discount
}

func CreateCoupon(name, store string, percent int64) IDiscount {
	return &Coupon{
		Discount: Discount{
			Name:    name,
			Percent: percent,
			Store:   store,
		},
	}
}
