/*
Algorithm: Pollard's p-1 for Integer Factorization

Input: n (the integer to be factorized), bLimit (the upper bound for the iteration)
Output: A non-trivial factor of n, or failure if no factor is found

1. Choose an initial value for a, typically a = 2
2. For b from 2 to bLimit:
    a. Calculate a = a^b mod n
3. Compute d = gcd(a - 1, n)

4. If d > 1 and d < n, return d (a non-trivial factor of n)
5. If no factor is found, return "failure"
*/

package main

import (
	"fmt"
	"math/big"
	"os"
)

// gcd calculates the Greatest Common Divisor of a and b.
func gcd(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a
	}
	return gcd(b, new(big.Int).Mod(a, b))
}

// pollardsPMinus1 is an implementation of Pollard's p-1 algorithm.
func pollardsPMinus1(n *big.Int, bLimit *big.Int) *big.Int {
	a := big.NewInt(2) // Initial value

	// Initialize b as a big.Int and start the loop
	b := big.NewInt(2)
	for b.Cmp(bLimit) < 0 {
		// a = a^b mod n
		a.Exp(a, b, n)

		// d = gcd(a-1, n)
		d := gcd(new(big.Int).Sub(a, big.NewInt(1)), n)

		if d.Cmp(big.NewInt(1)) > 0 && d.Cmp(n) < 0 {
			return d
		}

		// Increment b
		b.Add(b, big.NewInt(1))
	}
	return nil // No factor found
}

// findFactors uses Pollard's p-1 algorithm to find a factor of n.
func findFactors(n *big.Int, bLimit *big.Int) (*big.Int, *big.Int) {
	factor1 := pollardsPMinus1(n, bLimit)
	if factor1 == nil {
		return nil, nil // No factor found
	}

	// Calculate the second factor
	factor2 := new(big.Int).Div(n, factor1)

	return factor1, factor2
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run pollards_p_minus_1.go [number]")
		os.Exit(1)
	}

	n := new(big.Int)
	n, ok := n.SetString(os.Args[1], 10)
	if !ok {
		fmt.Println("Invalid number")
		os.Exit(1)
	}

	bitLen := big.NewInt(int64(n.BitLen()))
	bLimit := new(big.Int).Exp(big.NewInt(2), bitLen, nil) // arbitrary limit
	fmt.Println("Arbitrary bLimit: ", bLimit)

	factor1, factor2 := findFactors(n, bLimit)
	if factor1 == nil || factor2 == nil {
		fmt.Println("No factors found using Pollard's p-1")
	} else {
		fmt.Println("Factors found:", factor1, "and", factor2)

		// Verification test
		test := new(big.Int).Mul(factor1, factor2)
		if test.Cmp(n) == 0 {
			fmt.Println("Verification test passed: factors are correct.")
		} else {
			fmt.Println("Verification test failed: factors are incorrect.")
		}
	}
}
