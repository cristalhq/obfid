package obfid

import (
	"errors"
	"math"
	"math/big"
)

// IsPrime reports whether number is a prime using big.Int.ProbablyPrime.
func IsPrime(prime uint64) bool {
	return big.NewInt(int64(prime)).ProbablyPrime(20)
}

var (
	ErrNotPrime = errors.New("number is not a prime")
	ErrTooLarge = errors.New("prime is larger than math.MaxInt64")
)

// Generator can encode and decode integers using prime numbers.
type Generator struct {
	prime   uint64
	inverse uint64
	random  uint64
}

// NewGenerator creates a new generator using the provided prime number and random.
func NewGenerator(prime, random uint64) (*Generator, error) {
	inverse, err := inverse(prime)
	if err != nil {
		return nil, err
	}

	generator := &Generator{
		prime:   prime,
		inverse: inverse,
		random:  random,
	}
	return generator, nil
}

const maxInt32 = uint64(math.MaxInt32)

// Encode returns obfuscated number.
func (g *Generator) Encode(number uint64) uint64 {
	return ((number * g.prime) & maxInt32) ^ g.random
}

// Decode returns the original (deobfuscated) number.
func (g *Generator) Decode(obfuscated uint64) uint64 {
	return ((obfuscated ^ g.random) * g.inverse) & maxInt32
}

// inverse calculates the inverse of prime.
func inverse(prime uint64) (uint64, error) {
	switch {
	case prime > math.MaxInt64:
		return 0, ErrTooLarge

	case !IsPrime(prime):
		return 0, ErrNotPrime

	default:
		// TODO: make math/big free
		p := big.NewInt(int64(prime))
		max := big.NewInt(int64(maxInt32 + 1))
		var res big.Int
		return res.ModInverse(p, max).Uint64(), nil
	}
}
