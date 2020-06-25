package credit

func Calculate(sum int64, period int64, rate int64) (monPayment int64, overpay int64, total int64) {

	monthRate := float64(rate) / 12 / 100
	part := AltPow((1 + monthRate), period)
	annuityK := (monthRate * part) / (part - 1) * 10000
	monPayment = int64(annuityK) * sum / 1000000
	overpay = monPayment*period - sum/100
	total = monPayment * period
	return
}

func AltPow(x float64, y int64) (z float64) {
	z = x
	for i := 0; int64(i) <= y; i++ {
		z = z * x
	}
	return
}
