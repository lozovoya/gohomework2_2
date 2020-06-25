package credit_test

import (
    "github.com/lozovoya/gohomework2_1/pkg/credit"
    "fmt"
)

func ExampleCalculate() {
    fmt.Println(credit.Calculate(1_000_000_00, 36, 20))
    // Output:
    // 35700 285200 1285200

}