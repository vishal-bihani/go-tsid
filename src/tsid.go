package main

const (
	TSID_EPOCH  int64 = 1672531200000 // 2023-01-01T00:00:00.000Z
	TSID_BYTES  int8  = 8
	TSID_CHARS  int8  = 13
	RANDOM_BITS int8  = 22
	RANDOM_MASK int32 = 0x003fffff
)

type Tsid struct {
	Number int64
}
