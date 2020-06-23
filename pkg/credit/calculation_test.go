package credit

import (
    "credit"
    "fmt"
)

func ExampleCalculate() {
    fmt.Println(credit.Calculate(1_000_000_00, 36, 20/12/100))
    // Output:
    // месячный платеж - 35700р, переплата - 285200р, общая выплата - 1285200р

}