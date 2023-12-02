package main

import (
	"errors"
	"sync"
	"time"
)

var lock = &sync.Mutex{}
var tsidFactoryInstance *tsidFactory

type tsidFactory struct {
	node        int32
	nodeBits    int32
	nodeMask    int32
	counter     int32
	counterBits int32
	counterMask int32
	lastTime    int64
	customEpoch int64
	time        time.Time
	random      Random
	randomBytes int32
}

func newTsidFactory(builder *tsidFactoryBuilder) (*tsidFactory, error) {
	tsidFactory := &tsidFactory{
		nodeBits:    builder.nodeBits,
		customEpoch: builder.GetCustomEpoch(),
		time:        builder.GetTime(),
		random:      builder.GetRandom(),
	}

	tsidFactory.counterBits = int32(RANDOM_BITS) - builder.nodeBits
	tsidFactory.counterMask = int32(RANDOM_MASK >> builder.nodeBits)
	tsidFactory.nodeMask = int32(RANDOM_MASK >> tsidFactory.counterBits)

	tsidFactory.randomBytes = ((tsidFactory.counterBits - 1) / 8) + 1

	tsidFactory.node = builder.node & int32(tsidFactory.nodeMask)
	tsidFactory.lastTime = tsidFactory.time.UnixMilli()
	randomNumber, err := tsidFactory.getRandomValue()
	if err != nil {
		return nil, err
	}

	tsidFactory.counter = randomNumber
	return tsidFactory, nil
}

func (factory *tsidFactory) Generate() (*tsid, error) {
	time, err := factory.getTime()
	if err != nil {
		return nil, err
	}

	sTime := time << RANDOM_BITS
	sNode := factory.node << factory.counterBits
	sCounter := factory.counter & factory.counterMask

	tsidNumber := int64(sTime | int64(sNode) | int64(sCounter))
	return NewTsid(tsidNumber), nil
}

func (factory *tsidFactory) getTime() (int64, error) {
	time := factory.time.UnixMilli()
	if time <= factory.lastTime {
		factory.counter++
		carry := factory.counter >> factory.counterBits
		factory.counter = factory.counter & factory.counterMask
		time = factory.lastTime + int64(carry)

	} else {
		value, err := factory.getRandomValue()
		if err != nil {
			return 0, err
		}
		factory.counter = value
	}
	factory.lastTime = time
	return (time - factory.customEpoch), nil
}

func (factory *tsidFactory) getRandomValue() (int32, error) {
	return factory.getRandomCounter()
}

func (factory *tsidFactory) getRandomCounter() (int32, error) {
	switch factory.random.(type) {
	case *byteRandom:
		{
			bytes, err := factory.random.NextBytes(factory.randomBytes)
			if err != nil {
				return 0, err
			}

			switch len(bytes) {
			case 1:
				return int32((bytes[0] & 0xff) & byte(factory.counterMask)), nil
			case 2:
				return ((int32(bytes[0]&0xff) << 8) | int32(bytes[1]&0xff)) & factory.counterMask, nil
			case 3:
				return ((int32(bytes[0]&0xff) << 16) | (int32(bytes[1]&0xff) << 8) |
					int32(bytes[2]&0xff)) & factory.counterMask, nil
			}
		}
	case *intRandom:
		{
			value, err := factory.random.NextInt()
			if err != nil {
				return 0, err
			}

			return int32(value & factory.counterMask), nil
		}
	}

	return 0, errors.New("invalid random")
}

type tsidFactoryBuilder struct {
	node        int32
	nodeBits    int32
	customEpoch int64
	time        time.Time
	random      Random
}

func TsidFactoryBuilder() *tsidFactoryBuilder {
	return &tsidFactoryBuilder{}
}

func (builder *tsidFactoryBuilder) WithNode(node int32) *tsidFactoryBuilder {
	builder.node = node
	return builder
}

func (builder *tsidFactoryBuilder) WithNodeBits(nodeBits int32) *tsidFactoryBuilder {
	builder.nodeBits = nodeBits
	return builder
}

func (builder *tsidFactoryBuilder) WithCustomEpoch(customEpoch int64) *tsidFactoryBuilder {
	builder.customEpoch = customEpoch
	return builder
}

func (builder *tsidFactoryBuilder) WithTime(time time.Time) *tsidFactoryBuilder {
	builder.time = time
	return builder
}

func (builder *tsidFactoryBuilder) WithRandom(random Random) *tsidFactoryBuilder {
	builder.random = random
	return builder
}

func (builder *tsidFactoryBuilder) GetTime() time.Time {
	if builder.time.IsZero() {
		builder.time = time.Now().UTC()
	}
	return builder.time
}

func (builder *tsidFactoryBuilder) GetRandom() Random {
	if builder.random == nil {
		randomSupplier := NewMathRandomSupplier()
		builder.random = NewIntRandom(randomSupplier)
	}

	return builder.random
}

func (builder *tsidFactoryBuilder) GetCustomEpoch() int64 {
	if builder.customEpoch == 0 {
		builder.customEpoch = TSID_EPOCH
	}
	return builder.customEpoch
}

func (builder *tsidFactoryBuilder) Build() (*tsidFactory, error) {
	if tsidFactoryInstance != nil {
		return tsidFactoryInstance, nil
	}

	lock.Lock()
	defer lock.Unlock()

	var err error = nil
	if tsidFactoryInstance == nil {
		tsidFactoryInstance, err = newTsidFactory(builder)
		if err != nil {
			return nil, err
		}
	}
	return tsidFactoryInstance, nil
}
