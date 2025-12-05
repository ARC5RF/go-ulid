package ulid

import (
	"crypto/rand"
	"math/big"

	"github.com/ARC5RF/go-blame"
)

func rand_int_64() (int64, error) {
	idx, rand_err := blame.O1(rand.Int(rand.Reader, big.NewInt(ENCODING_LEN)))
	if rand_err != nil {
		return 0, rand_err
	}
	return idx.Int64(), nil
}

func random_char(rng PRNG) (string, error) {
	want, rand_err := blame.O1(rng())
	if rand_err != nil {
		return "", rand_err
	}
	return ENCODING[want : want+1], nil
}

func replace_char_at(input string, index int, with byte) *string {
	if index > len(input)-1 {
		return &input
	}

	o := input[0:index] + string(with) + input[index+1:]
	return &o
}
