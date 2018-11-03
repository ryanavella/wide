// Package wide provides implementations of Int128 and Uint128 for Go. See readme.md for more information.
package wide

// Data sizes
const (
	int32Size  = 32
	int64Size  = 64
	int128Size = 128
)

// Maximum and minimum integer sizes
const (
	maxInt64  = 1<<63 - 1
	minInt64  = -1<<63
	maxUint64 = 1<<64 - 1
)
