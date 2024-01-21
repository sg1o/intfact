/*
FermatFactorization(N):

	// Step 1: Initialization
	a = ceil(sqrt(N))
	b2 = a*a - N

	// Step 2: Search for factors
	while not isSquare(b2):
	    a = a + 1
	    b2 = a*a - N

	// Step 3: Calculate b and factors
	b = sqrt(b2)
	factor1 = a + b
	factor2 = a - b

	return (factor1, factor2)

isSquare(x):

	// Check if x is a perfect square
	root = floor(sqrt(x))
	return root*root == x

// Usage Example
N = Product of two twin primes (e.g., 11 * 13)
factors = FermatFactorization(N)
print("The factors are:", factors)
*/
package main

import (
	"fmt"
	"math/big"
	"os"
)

// fermatFactorization attempts to factorize a number using Fermat's factorization method.
func fermatFactorization(n *big.Int) (*big.Int, *big.Int) {
	a := big.NewInt(0).Set(n)
	a.Sqrt(a).Add(a, big.NewInt(1))

	bSquared := big.NewInt(0).Mul(a, a)
	bSquared.Sub(bSquared, n)

	b := new(big.Int)

	for bSquared.Sign() >= 0 {
		b.Sqrt(bSquared)
		if bSquared.Cmp(n) == 0 {
			break
		}
		a.Add(a, big.NewInt(1))
		bSquared.Mul(a, a).Sub(bSquared, n)
	}

	if bSquared.Sign() < 0 {
		return nil, nil // No factors found
	}

	p := new(big.Int).Sub(a, b)
	q := new(big.Int).Add(a, b)

	return p, q // Factors found
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run fermat.go [number]")
		os.Exit(1)
	}

	n := new(big.Int)
	n, ok := n.SetString(os.Args[1], 10)
	if !ok {
		fmt.Println("Invalid number")
		os.Exit(1)
	}

	p, q := fermatFactorization(n)
	if p == nil || q == nil {
		fmt.Println("No factors found using Fermat's factorization")
	} else {
		fmt.Println("Factors found:", p, q)
	}
}
