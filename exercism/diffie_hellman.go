// Package diffiehellman provides a Go solution for the Exercism diffie-hellman exercise
package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

var randy = rand.New(rand.NewSource(time.Now().UnixNano()))

// PrivateKey returns a private key - a random number in the range [2, p)
func PrivateKey(p *big.Int) *big.Int {
	acc := big.NewInt(0)
	two := big.NewInt(2)

	acc.Set(p)
	acc.Sub(acc.Set(p), two)
	acc.Rand(randy, acc)
	acc.Add(acc, two)

	return acc
}

// PublicKey returns a public key given a private key
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return key(p, big.NewInt(g), private)
}

// SecretKey returns a secret given a private key (a) and a public key (b_pub)
func SecretKey(private, public, p *big.Int) *big.Int {
	return key(p, public, private)
}

// NewPair return a private and public key
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)

	public = PublicKey(private, p, g)

	return
}

// key does the diffie-hellman bit
func key(modulo, number, power *big.Int) *big.Int {
	acc := big.NewInt(0)

	return acc.Exp(acc.Set(number), power, modulo)
}
