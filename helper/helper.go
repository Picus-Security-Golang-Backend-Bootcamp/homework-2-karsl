package helper

import (
	"crypto/rand"
	"math"
	"math/big"
)

func GetRandomFloat64(max, precision int) (float64, error) {
	randInt, err := GetRandomInt64(int64(float64(max) * math.Pow10(precision)))
	if err != nil {
		return 0, nil
	}

	return float64(randInt) / math.Pow10(precision), nil
}

func GetRandomInt64(max int64) (int64, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, nil
	}

	return nBig.Int64(), nil
}
