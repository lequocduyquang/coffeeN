package main

import "fmt"

func NewDiscount(source, name, store string, percent int64) (IDiscount, error) {
	if source == "voucher" {
		return CreateVoucher(name, store, percent), nil
	}
	if source == "coupon" {
		return CreateCoupon(name, store, percent), nil
	}
	return nil, fmt.Errorf("Not match source discount")
}
