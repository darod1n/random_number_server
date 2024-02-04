package ports

import "math/big"

type GeneratorRandomNumber interface {
	RandNumber() *big.Int
}
