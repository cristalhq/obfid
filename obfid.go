package obfid

import (
	"errors"
	"math"
	"math/big"
)

// IsPrime reports whether number is a prime using big.Int.ProbablyPrime.
func IsPrime(prime uint64) bool {
	return (&big.Int{}).SetUint64(prime).ProbablyPrime(20)
}

var (
	ErrTooMuchBits = errors.New("too much bits")
	ErrNotPrime    = errors.New("number is not a prime")
	ErrTooLarge    = errors.New("prime is larger than math.MaxInt64")
)

// Generator can encode and decode integers using prime numbers.
type Generator struct {
	prime   uint64
	inverse uint64
	random  uint64
	mask    uint64
}

// NewGenerator creates a new generator using the provided prime number and random.
func NewGenerator(prime, random uint64, bits int) (*Generator, error) {
	if bits == 0 || bits > 64 {
		return nil, ErrTooMuchBits
	}
	mask := uint64(1<<bits - 1)

	inverse, err := inverse(prime, mask)
	if err != nil {
		return nil, err
	}

	generator := &Generator{
		prime:   prime,
		inverse: inverse,
		random:  random,
		mask:    mask,
	}
	return generator, nil
}

// Encode returns obfuscated number.
func (g *Generator) Encode(number uint64) uint64 {
	return ((number * g.prime) & g.mask) ^ g.random
}

// Decode returns the original (deobfuscated) number.
func (g *Generator) Decode(obfuscated uint64) uint64 {
	return ((obfuscated ^ g.random) * g.inverse) & g.mask
}

// inverse calculates the inverse of prime.
func inverse(prime, mask uint64) (uint64, error) {
	switch {
	case prime > math.MaxInt64:
		return 0, ErrTooLarge

	case !IsPrime(prime):
		return 0, ErrNotPrime

	default:
		// TODO: make alloc free (without math/big pkg)
		p := (&big.Int{}).SetUint64(prime)
		max := (&big.Int{}).SetUint64(mask + 1)
		return (&big.Int{}).ModInverse(p, max).Uint64(), nil
	}
}
