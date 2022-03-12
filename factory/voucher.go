package main

type Voucher struct {
	Discount
}

func CreateVoucher(name, store string, percent int64) IDiscount {
	return &Voucher{
		Discount: Discount{
			Name:    name,
			Percent: percent,
			Store:   store,
		},
	}
}
