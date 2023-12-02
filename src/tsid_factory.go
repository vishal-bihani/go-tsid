package main

import (
	"sync"
	"time"
)

var lock = &sync.Mutex{}
var tsidFactoryInstance *tsidFactory

type tsidFactory struct {
	node        int32
	nodeBits    int32
	nodeMask    uint32
	counter     int32
	counterBits int32
	counterMask uint32
	lastTime    int64
	customEpoch int64
	time        *time.Time
	random      Random
}

func newTsidFactory(builder *tsidFactoryBuilder) (*tsidFactory, error) {
	tsidFactory := &tsidFactory{
		nodeBits:    builder.nodeBits,
		customEpoch: builder.customEpoch,
		time:        builder.time,
		random:      builder.random,
	}

	tsidFactory.counterBits = int32(RANDOM_BITS) - builder.nodeBits
	tsidFactory.counterMask = uint32(RANDOM_MASK >> builder.nodeBits)
	tsidFactory.nodeMask = uint32(RANDOM_MASK >> tsidFactory.counterBits)

	tsidFactory.node = builder.node & int32(tsidFactory.nodeMask)
	tsidFactory.lastTime = tsidFactory.time.UnixMilli()
	randomNumber, err := tsidFactory.random.NextInt()
	if err != nil {
		return nil, err
	}

	tsidFactory.counter = randomNumber
	return tsidFactory, nil
}

type tsidFactoryBuilder struct {
	node        int32
	nodeBits    int32
	customEpoch int64
	time        *time.Time
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

func (builder *tsidFactoryBuilder) WithTime(time *time.Time) *tsidFactoryBuilder {
	builder.time = time
	return builder
}

func (builder *tsidFactoryBuilder) WithRandom(random Random) *tsidFactoryBuilder {
	builder.random = random
	return builder
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
