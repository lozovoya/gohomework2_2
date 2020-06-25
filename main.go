package main

import (
	"fmt"
	"github.com/lozovoya/gohomework2_2/pkg/credit"
)

func main() {

	var sum int64 = 1_000_000_00
	var period int64 = 36
	var rate int64 = 20

	mon_payment, overpay, total := credit.Calculate(sum, period, rate)
	fmt.Printf("месячный платеж - %vр, переплата - %vр, общая выплата - %vр\n", mon_payment, overpay, total)

}
