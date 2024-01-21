/*
Algorithm: Pollard's Rho for Integer Factorization

Input: n (the integer to factorize)
Output: A non-trivial factor of n, or failure if no factor is found

1. Choose an initial value for x, typically x = 2
2. Set y = x
3. Choose a function f, commonly f(x) = (x^2 + 1) mod n
4. Set d = 1 (to store the gcd)

5. While d == 1:
    a. Set x = f(x)
    b. Set y = f(f(y))
    c. Set d = gcd(|x - y|, n)

6. If d == n, return "failure" (the algorithm did not find a factor)
7. Else, return d (a non-trivial factor of n)
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

// f is the polynomial function used in Pollard's rho: f(x) = x^2 + 1.
func f(x, n *big.Int) *big.Int {
	return new(big.Int).Mod(new(big.Int).Add(new(big.Int).Mul(x, x), big.NewInt(1)), n)
}

// pollardsRho is an implementation of Pollard's Rho algorithm.
func pollardsRho(n *big.Int) *big.Int {
	x := big.NewInt(2) // Initial value of x
	y := big.NewInt(2) // Initial value of y
	d := big.NewInt(1) // gcd(x-y, n)

	for d.Cmp(big.NewInt(1)) == 0 {
		x = f(x, n)       // x = f(x)
		y = f(f(y, n), n) // y = f(f(y))
		d = gcd(new(big.Int).Abs(new(big.Int).Sub(x, y)), n)
	}

	if d.Cmp(n) == 0 {
		return nil // Failure, a non-trivial factor was not found
	}
	return d // Non-trivial factor found
}

// findFactors uses Pollard's Rho algorithm to find two factors of n.
func findFactors(n *big.Int) (*big.Int, *big.Int) {
	factor1 := pollardsRho(n)
	if factor1 == nil {
		return nil, nil // No factor found
	}

	// Calculate the second factor
	factor2 := new(big.Int).Div(n, factor1)

	return factor1, factor2
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run pollards_rho.go [number]")
		os.Exit(1)
	}

	n := new(big.Int)
	n, ok := n.SetString(os.Args[1], 10)
	if !ok {
		fmt.Println("Invalid number")
		os.Exit(1)
	}

	factor1, factor2 := findFactors(n)
	if factor1 == nil || factor2 == nil {
		fmt.Println("No factors found using Pollard's Rho")
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
