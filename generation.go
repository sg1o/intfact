package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// generatePrime generates a random prime number of specified bit size.
func generatePrime(bitSize int) *big.Int {
	prime, err := rand.Prime(rand.Reader, bitSize)
	if err != nil {
		fmt.Println("Error generating prime:", err)
		return nil
	}
	return prime
}

// generateSafePrime generates a safe prime number.
func generateSafePrime(bitSize int) *big.Int {
	for {
		prime := generatePrime(bitSize)
		safePrimeCandidate := new(big.Int).Sub(prime, big.NewInt(1))
		safePrimeCandidate.Div(safePrimeCandidate, big.NewInt(2))
		if safePrimeCandidate.ProbablyPrime(20) { // primality test with 20 rounds
			return prime
		}
	}
}

// avoidPrimeTwin generates a prime number that is not a twin prime.
func avoidPrimeTwin(basePrime *big.Int) *big.Int {
	for {
		newPrime := generateSafePrime(basePrime.BitLen())
		if new(big.Int).Abs(new(big.Int).Sub(basePrime, newPrime)).Cmp(big.NewInt(2)) != 0 {
			return newPrime
		}
	}
}

// generateFirstNPrimes generates the first N prime numbers.
func generateFirstNPrimes(n int) []int {
	primes := []int{}
	for i := 2; len(primes) < n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// isPrime checks if a number is prime (simple trial division method).
func isPrime(number int) bool {
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// hasSmallFactors checks if a number has small prime factors from the first 100 primes.
func hasSmallFactors(number *big.Int) bool {
	smallPrimes := generateFirstNPrimes(100)
	for _, prime := range smallPrimes {
		if new(big.Int).Mod(number, big.NewInt(int64(prime))).Cmp(big.NewInt(0)) == 0 {
			return true
		}
	}
	return false
}

// generateTwinPrimes generates a pair of twin primes.
func generateTwinPrimes(bitSize int) (*big.Int, *big.Int) {
	for {
		p := generatePrime(bitSize)
		twinCandidate := new(big.Int).Add(p, big.NewInt(2))
		if twinCandidate.ProbablyPrime(20) {
			return p, twinCandidate
		}
	}
}

// generatePMinusOneComposite generates a composite number where one prime's p-1 has small factors.
func generatePMinusOneComposite(bitSize int) *big.Int {
	for {
		p := generatePrime(bitSize)
		pMinusOne := new(big.Int).Sub(p, big.NewInt(1))

		if hasSmallFactors(pMinusOne) {
			q := generatePrime(bitSize) // Generate another prime for the composite
			return new(big.Int).Mul(p, q)
		}
	}
}

// generateTwinPrimeComposite generates a composite number from twin primes.
func generateTwinPrimeComposite(bitSize int) *big.Int {
	p, q := generateTwinPrimes(bitSize)
	return new(big.Int).Mul(p, q)
}

// generateClosePrime generates a prime number close to a given prime.
func generateClosePrime(basePrime *big.Int) *big.Int {
	offset := big.NewInt(2) // Start with an offset of 2
	closePrime := new(big.Int).Add(basePrime, offset)

	for !closePrime.ProbablyPrime(20) { // Using Miller-Rabin test, enough for our purpose
		offset.Add(offset, big.NewInt(2)) // Try the next odd number
		closePrime.Add(basePrime, offset)
	}
	return closePrime
}

func generateCloseComposite(bitSize int) *big.Int {
	p := generatePrime(bitSize)
	q := generateClosePrime(p)

	return new(big.Int).Mul(p, q)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run generation.go <bitSize>")
		os.Exit(1)
	}

	bitSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid bit size:", err)
		os.Exit(1)
	}

	// Generate normal composite number
	p := generatePrime(bitSize)
	q := generatePrime(bitSize)
	normalComposite := new(big.Int).Mul(p, q)
	fmt.Println("Normal composite number:", normalComposite)

	// Generate strong composite number
	strongComposite := new(big.Int).Mul(generateSafePrime(bitSize), avoidPrimeTwin(p))
	fmt.Println("Strong composite number:", strongComposite)

	// Generate P-1 composite number
	pMinusOneComposite := generatePMinusOneComposite(bitSize)
	fmt.Println("P-1 composite number:", pMinusOneComposite)

	// Generate twin prime composite number
	twinPrimeComposite := generateTwinPrimeComposite(bitSize)
	fmt.Println("Twin prime composite number:", twinPrimeComposite)

	// Close primes
	closeComposition := generateCloseComposite(bitSize)
	fmt.Println("Close primes composition:", closeComposition)
}
