package tsid

type Clock interface {
	UnixMilli() int64
}
