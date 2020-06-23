package main

import (
	"fmt"
	"github.com/lozovoya/gohomework2_1/pkg/credit"
)

func main() {

	var sum int64 = 1_000_000_00
	var period int64 = 36
	var rate float64 = 20
	monthRate := rate / 12 / 100

	mon_payment, overpay, total := credit.Calculate(sum, period, monthRate)
	fmt.Printf("месячный платеж - %vр, переплата - %vр, общая выплата - %vр\n", mon_payment, overpay, total)

}
