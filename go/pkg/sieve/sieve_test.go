package sieve

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthPrime(t *testing.T) {
	sieve := NewSieve()

	// I was getting an error with the tests and the assert package below.
	// I couldn't tell if it was just something on my end or in my IDE,
	// but it was taking more time to try and fix it than I was anticipating.  So I add these
	// additional tests that I was able to run and validate my code.  Apologies if this was intentional
	// and I was supposed to fix them :)
	// if got := sieve.NthPrime(0); got != int64(2) {
	// 	t.Errorf("NthPrime(0) = %v; want %v", got, int64(2))
	// }
	// if got := sieve.NthPrime(19); got != int64(71) {
	// 	t.Errorf("NthPrime(19) = %v; want %v", got, int64(71))
	// }
	// if got := sieve.NthPrime(99); got != int64(541) {
	// 	t.Errorf("NthPrime(99) = %v; want %v", got, int64(541))
	// }
	// if got := sieve.NthPrime(500); got != int64(3581) {
	// 	t.Errorf("NthPrime(500) = %v; want %v", got, int64(3581))
	// }
	// if got := sieve.NthPrime(986); got != int64(7793) {
	// 	t.Errorf("NthPrime(986) = %v; want %v", got, int64(7793))
	// }
	// if got := sieve.NthPrime(2000); got != int64(17393) {
	// 	t.Errorf("NthPrime(2000) = %v; want %v", got, int64(17393))
	// }
	// if got := sieve.NthPrime(1000000); got != int64(15485867) {
	// 	t.Errorf("NthPrime(1000000) = %v; want %v", got, int64(15485867))
	// }
	// if got := sieve.NthPrime(10000000); got != int64(179424691) {
	// 	t.Errorf("NthPrime(10000000) = %v; want %v", got, int64(179424691))
	// }

	assert.Equal(t, int64(2), sieve.NthPrime(0))
	assert.Equal(t, int64(71), sieve.NthPrime(19))
	assert.Equal(t, int64(541), sieve.NthPrime(99))
	assert.Equal(t, int64(3581), sieve.NthPrime(500))
	assert.Equal(t, int64(7793), sieve.NthPrime(986))
	assert.Equal(t, int64(17393), sieve.NthPrime(2000))
	assert.Equal(t, int64(15485867), sieve.NthPrime(1000000))
	assert.Equal(t, int64(179424691), sieve.NthPrime(10000000))
	assert.Equal(t, int64(2038074751), sieve.NthPrime(100000000))
}

func FuzzNthPrime(f *testing.F) {
	sieve := NewSieve()

	f.Fuzz(func(t *testing.T, n int64) {
		if !big.NewInt(sieve.NthPrime(n)).ProbablyPrime(0) {
			t.Errorf("the sieve produced a non-prime number at index %d", n)
		}
	})
}
