package main

import (
	"fmt"
	"math"
)

func snowball(principal float64, interestRate float64, numberOfPayments int) float64 {
	r := interestRate / 100.0 / 12.0
	return principal * (math.Pow(1 + r, float64(numberOfPayments)) - 1) / (r * math.Pow(1 + r, float64(numberOfPayments)))
}

func main() {
	result := snowball(1000.0, 10.0, 36)
	fmt.Println("Monthly payment:", result)
}
