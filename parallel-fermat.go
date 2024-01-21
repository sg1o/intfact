package main

import (
	"fmt"
	"math/big"
	"os"
	"runtime"
	"sync"
)

type FactorPair struct {
	p, q *big.Int
}

func parallelFermatFactorization(n, start, end *big.Int, wg *sync.WaitGroup, result chan<- FactorPair) {
	defer wg.Done()

	a := big.NewInt(0).Set(start)
	bSquared := big.NewInt(0).Mul(a, a)
	bSquared.Sub(bSquared, n)
	b := new(big.Int)

	for a.Cmp(end) <= 0 {
		b.Mul(b, b)
		if bSquared.Cmp(b) == 0 {
			result <- FactorPair{new(big.Int).Sub(a, b), new(big.Int).Add(a, b)}
			return
		}
		a.Add(a, big.NewInt(1))
		bSquared.Mul(a, a).Sub(bSquared, n)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run parallel-fermat.go [number]")
		os.Exit(1)
	}

	n := new(big.Int)
	_, ok := n.SetString(os.Args[1], 10)
	if !ok {
		fmt.Println("Invalid number")
		os.Exit(1)
	}

	numThreads := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(numThreads)
	result := make(chan FactorPair)

	start := new(big.Int).Sqrt(n)
	increment := new(big.Int).Sub(new(big.Int).Add(start, big.NewInt(1)), start)
	increment.Div(increment, big.NewInt(int64(numThreads)))

	for i := 0; i < numThreads; i++ {
		end := new(big.Int).Add(start, increment)
		if i == numThreads-1 {
			end.Set(new(big.Int).Add(n, big.NewInt(1)))
		}
		go parallelFermatFactorization(n, start, end, &wg, result)
		start = new(big.Int).Add(end, big.NewInt(1))
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	r, ok := <-result
	if ok {
		fmt.Println("Factors found:", r.p, r.q)
	} else {
		fmt.Println("No factors found using Fermat's factorization")
	}
}
