// Tax Office Problem
// ------------------
// Given an annual income p (1000 ≤ p ≤ 10^8), calculate the tax owed.
// The tax is computed in a progressive (tiered) system:
//
// - First 100 coins: taxed at 5%
// - Next 400 coins: taxed at 10%
// - Next 500 coins: taxed at 15%
// - Any remaining income: taxed at 20%
//
// The result must be rounded down to the nearest integer.
//
// Example:
// Income = 2000
// Tax = (100 * 0.05) + (400 * 0.10) + (500 * 0.15) + (1000 * 0.20)
// Tax = 5 + 40 + 75 + 200 = 320
//
// Input: integer p (annual income)
// Output: integer (total tax)


package main


import "fmt"

func main() {
	var income, tax int

	fmt.Scanf("%d", &income)

	switch {
		
		case income <= 100:
			tax += income * 5 / 100

		case income <= 500:
			tax += 5
			tax += (income - 100) / 10

		case income <= 1000:
			tax += 45
			tax += (income - 500) * 15 / 100

		default:
			tax += 90
			tax += (income - 1000) / 5

	}
	
	fmt.Println(tax)
}