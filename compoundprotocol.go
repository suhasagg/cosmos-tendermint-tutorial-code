package main

import (
	"fmt"
)

// Define a struct to represent the Compound protocol
type CompoundProtocol struct {
	// Define fields that represent the properties of the protocol
	balance float64
	borrow float64
	interest float64
	liquidationPrice float64
}

// Define a method to deposit funds into the Compound protocol
func (c *CompoundProtocol) Deposit(amount float64) {
	c.balance += amount
}

// Define a method to withdraw funds from the Compound protocol
func (c *CompoundProtocol) Withdraw(amount float64) {
	c.balance -= amount
}

// Define a method to calculate the interest earned
func (c *CompoundProtocol) CalculateInterest() {
	c.interest = c.balance * 0.05
}

// Define a method to check the liquidation price
func (c *CompoundProtocol) CheckLiquidationPrice() {
	c.liquidationPrice = c.borrow * 2
}

func main() {
	// Create an instance of the CompoundProtocol struct
	compound := CompoundProtocol{
		balance: 1000.0,
		borrow: 500.0,
		interest: 0.0,
		liquidationPrice: 0.0,
	}

	// Deposit funds into the Compound protocol
	compound.Deposit(500.0)

	// Withdraw funds from the Compound protocol
	compound.Withdraw(200.0)

	// Calculate the interest earned
	compound.CalculateInterest()

	// Check the liquidation price
	compound.CheckLiquidationPrice()

	// Print the final values of the Compound protocol
	fmt.Println("Balance: ", compound.balance)
	fmt.Println("Borrow: ", compound.borrow)
	fmt.Println("Interest: ", compound.interest)
	fmt.Println("Liquidation Price: ", compound.liquidationPrice)
}
