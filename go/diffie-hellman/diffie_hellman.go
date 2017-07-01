package diffiehellman

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"time"
)

const testVersion = 1

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func PrivateKey(p *big.Int) *big.Int {
	n, _ := crand.Int(rnd, p)
	for n.Int64() <= 1 || n.Cmp(p) >= 0 {
		n, _ = crand.Int(rnd, p)
	}
	return n
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	publicKey := big.NewInt(0)
	return publicKey.Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	secret := big.NewInt(0).Exp(public2, private1, p)
	return secret
}
