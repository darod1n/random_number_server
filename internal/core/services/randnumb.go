package service

import (
	"context"
	"crypto/rand"
	"math/big"
)

type RandomNumberRepository interface {
	Exist(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key string, value string) error
}

type GeneratorRandomNumber struct {
	repository RandomNumberRepository
	max        int64
}

func (grs *GeneratorRandomNumber) RandNumber() *big.Int {
	for {
		randNum, err := rand.Int(rand.Reader, big.NewInt(grs.max))
		if err != nil {
			continue
		}
		ok, err := grs.repository.Exist(context.Background(), randNum.String())
		if err != nil {
			continue
		}
		if ok {
			continue
		}
		return randNum

	}
}

func NewGeneratorRandomNumber(
	repository RandomNumberRepository,
	max int64,
) *GeneratorRandomNumber {
	return &GeneratorRandomNumber{
		repository: repository,
		max:        max,
	}
}
