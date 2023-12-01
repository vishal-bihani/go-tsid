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
}

func newTsidFactory(builder *tsidFactoryBuilder) *tsidFactory {
	tsidFactory := &tsidFactory{
		nodeBits:    builder.nodeBits,
		customEpoch: builder.customEpoch,
		time:        builder.time,
	}

	tsidFactory.counterBits = int32(RANDOM_BITS) - builder.nodeBits
	tsidFactory.counterMask = uint32(RANDOM_MASK >> builder.nodeBits)
	tsidFactory.nodeMask = uint32(RANDOM_MASK >> tsidFactory.counterBits)

	tsidFactory.node = builder.node & int32(tsidFactory.nodeMask)
	tsidFactory.lastTime = tsidFactory.time.UnixMilli()

	return tsidFactory
}

type tsidFactoryBuilder struct {
	node        int32
	nodeBits    int32
	customEpoch int64
	time        *time.Time
}

func TsidFactoryBuilder() *tsidFactoryBuilder {
	return &tsidFactoryBuilder{}
}

func (builder *tsidFactoryBuilder) WithNode(node int32) {
	builder.node = node
}

func (builder *tsidFactoryBuilder) WithNodeBits(nodeBits int32) {
	builder.nodeBits = nodeBits
}

func (builder *tsidFactoryBuilder) WithCustomEpoch(customEpoch int64) {
	builder.customEpoch = customEpoch
}

func (builder *tsidFactoryBuilder) WithTime(time *time.Time) {
	builder.time = time
}

func (builder *tsidFactoryBuilder) Build() *tsidFactory {
	if tsidFactoryInstance != nil {
		return tsidFactoryInstance
	}

	lock.Lock()
	defer lock.Unlock()

	if tsidFactoryInstance == nil {
		tsidFactoryInstance = newTsidFactory(builder)
	}
	return tsidFactoryInstance
}
