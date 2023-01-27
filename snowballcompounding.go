package main

import (
    "fmt"
    "math/big"
)

func main() {
    // Initial stake
    stake := big.NewInt(100)
    // Interest rate
    interestRate := 0.05
    // Number of compounding periods
    periods := 12

    // Snowball compounding loop
    for i := 0; i < periods; i++ {
        // Calculate interest
        interest := new(big.Int).Mul(stake, big.NewInt(int64(interestRate*100)))
        interest.Quo(interest, big.NewInt(100))
        // Add interest to stake
        stake.Add(stake, interest)
        // Print current stake value
        fmt.Printf("After period %d: %s\n", i+1, stake)
    }
}
