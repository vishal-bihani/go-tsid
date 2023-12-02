package main

import (
	crypto_rand "crypto/rand"
	"math"
	"math/big"
	math_rand "math/rand"
	"time"
)

const (
	BYTE_SIZE        = 8
	INTEGER_SIZE_32  = 32
	INTEGER_BYTES_32 = 4
)

type Random interface {
	NextInt() (int32, error)
	NextBytes(length int32) ([]byte, error)
}

type intRandom struct {
	randomSupplier RandomSupplier
}

func NewIntRandom(randomSupplier RandomSupplier) *intRandom {
	return &intRandom{
		randomSupplier: randomSupplier,
	}
}

func (i *intRandom) NextInt() (int32, error) {
	return i.randomSupplier.GetInt()
}

func (i *intRandom) NextBytes(length int32) ([]byte, error) {
	bytes := make([]byte, length)

	shift := 0
	var random int32 = 0
	var err error = nil

	for j := 0; j < int(length); j++ {

		if shift < BYTE_SIZE {
			shift = INTEGER_SIZE_32
			random, err = i.randomSupplier.GetInt()
			if err != nil {
				return nil, err
			}
		}
		shift -= BYTE_SIZE
		bytes[j] = byte(uint32(random >> shift))
	}

	return bytes, nil
}

type byteRandom struct {
	randomSupplier RandomSupplier
}

func NewByteRandom(randomSupplier RandomSupplier) *byteRandom {
	return &byteRandom{
		randomSupplier: randomSupplier,
	}
}

func (i *byteRandom) NextInt() (int32, error) {
	var number int32 = 0
	bytes, err := i.randomSupplier.GetBytes(INTEGER_SIZE_32)
	if err != nil {
		return int32(number), err
	}

	for j := 0; j < INTEGER_BYTES_32; j++ {
		number = int32(byte(number<<BYTE_SIZE) | (bytes[j] & 0xff))
	}
	return number, nil
}

func (i *byteRandom) NextBytes(length int32) ([]byte, error) {
	return i.randomSupplier.GetBytes(length)
}

// Suppliers
type RandomSupplier interface {
	GetInt() (int32, error)
	GetBytes(length int32) ([]byte, error)
}

type mathRandom struct {
	rand *math_rand.Rand
}

func NewMathRandom() *mathRandom {
	return &mathRandom{
		rand: &math_rand.Rand{},
	}
}

func (i *mathRandom) GetInt() (int32, error) {
	i.rand.Seed(
		time.Now().
			UnixNano())

	return i.rand.Int31(), nil
}

func (i *mathRandom) GetBytes(length int32) ([]byte, error) {

	bytes := make([]byte, length)
	_, err := i.rand.Read(bytes)

	return bytes, err
}

type cryptoRandom struct {
}

func NewCryptoRandom() *cryptoRandom {
	return &cryptoRandom{}
}

func (i *cryptoRandom) GetInt() (int32, error) {
	random, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(math.MaxInt32))
	return int32(random.Int64()), err
}

func (i *cryptoRandom) GetBytes(length int32) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := crypto_rand.Read(bytes)

	return bytes, err
}
