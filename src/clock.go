package main

type Clock interface {
	UnixMilli() int64
}
