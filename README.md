# Prime Factorization Project

This project contains a collection of Go programs aimed at performing prime
factorization using various algorithms. Each file in this project implements a
different factorization method or related utility.

## File Descriptions

- **generation.go**: This script is likely used for generating large numbers or
  primes that can be used for testing the factorization algorithms.

- **fermat.go**: Implements Fermat's Factorization Method. This method is
  particularly effective for integers that are a product of two primes that are
  close to each other, such as twin primes.

- **pollards_rho.go**: Implements Pollard's Rho algorithm. This is a
  probabilistic factorization algorithm that is effective for large numbers.

- **parallel-fermat.go**: A parallelized version of Fermat's Factorization
  Method. It utilizes multiple CPU cores to divide and conquer the
  factorization problem, potentially leading to faster results for large
  numbers.

- **pollards_p_minus_1.go**: Implements Pollard's p-1 algorithm. This algorithm
  is effective in factorizing numbers with small prime factors.

## Usage

To use these scripts, ensure you have Go installed on your system. You can run
each script from the command line, passing the number to be factorized as an
argument. For example:

```bash
go run fermat.go 123456789
```

Replace `fermat.go` with the name of the script you wish to run and `123456789`
with the number you want to factorize.

## Requirements

- Go Programming Language (version 1.x or later)
- Basic understanding of command-line operations

## Contributing

Contributions to this project are welcome. Please ensure that any pull requests
or issues are descriptive and relevant to the project's goals.
