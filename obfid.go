package obfid

import (
	"errors"
	"math/big"
)

// IsPrime reports whether number is a prime using big.Int.ProbablyPrime.
func IsPrime(prime uint64) bool {
	return (&big.Int{}).SetUint64(prime).ProbablyPrime(20)
}

var (
	ErrInvalidBits = errors.New("too much bits")
	ErrTooMuchBits = ErrInvalidBits
	ErrEvenPrime   = errors.New("prime cannot be even")
	ErrNotPrime    = errors.New("number is not a prime")
	ErrTooLarge    = errors.New("prime is larger than math.MaxInt64")
)

// Generator can encode and decode integers using prime numbers.
type Generator struct {
	prime   uint64
	inverse uint64
	xorMask uint64
	offset  uint64
	mask    uint64
}

// NewGenerator creates a new generator using the provided prime number, random and offset.
func NewGenerator(prime, xorMask, offset uint64, bits int) (*Generator, error) {
	switch {
	case bits < 1 || bits > 64:
		return nil, ErrInvalidBits
	case prime%2 == 0:
		return nil, ErrEvenPrime
	case !IsPrime(prime):
		return nil, ErrNotPrime
	}

	mask := uint64(1<<bits - 1)

	generator := &Generator{
		prime:   prime,
		inverse: modInverse(prime) & mask,
		xorMask: xorMask,
		offset:  offset,
		mask:    mask,
	}
	return generator, nil
}

// Encode returns obfuscated number.
func (g *Generator) Encode(number uint64) uint64 {
	return (((number * g.prime) & g.mask) ^ g.xorMask) + g.offset
}

// Decode returns the original (deobfuscated) number.
func (g *Generator) Decode(obfuscated uint64) uint64 {
	return (((obfuscated - g.offset) ^ g.xorMask) * g.inverse) & g.mask
}

// modInverse returns the inverse of odd p modulo 2^64.
func modInverse(p uint64) uint64 {
	inv := p                 // correct to 3 bits, since p*p ≡ 1 (mod 8)
	for i := 0; i < 5; i++ { // 3 → 6 → 12 → 24 → 48 → 96 bits
		inv *= 2 - p*inv
	}
	return inv
}
