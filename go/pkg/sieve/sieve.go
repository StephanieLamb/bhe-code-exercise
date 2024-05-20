package sieve

import (
	"math"
)

type Sieve interface {
	NthPrime(n int64) int64
}

// PrimeSieve struct to implement the Sieve interface
type PrimeSieve struct{}

func NewSieve() Sieve {
	return &PrimeSieve{}
}

func (s *PrimeSieve) SieveOfEratosthenes(limit int64) []int64 {
	if limit < 2 {
		return []int64{}
	}

	// initialize a boolean array to mark prime numbers
	isPrime := make([]bool, limit+1)
	for i := int64(2); i <= limit; i++ {
		isPrime[i] = true
	}

	// sieve of Eratosthenes
	for p := int64(2); p*p <= limit; p++ {
		if isPrime[p] {
			for multiple := p * p; multiple <= limit; multiple += p {
				isPrime[multiple] = false
			}
		}
	}

	// add prime numbers to slice
	primes := []int64{}
	for p := int64(2); p <= limit; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

func (s *PrimeSieve) NthPrime(n int64) int64 {
	if n < 0 {
		panic("n cannot be a negative number")
	}

	scope := int64(float64(n) * math.Log(float64(n)))
	if scope < 2 {
		scope = 2
	}

	primes := s.SieveOfEratosthenes(scope)

	for int64(len(primes)) <= n {
		scope *= 2
		primes = s.SieveOfEratosthenes(scope)
	}

	return primes[n]
}
