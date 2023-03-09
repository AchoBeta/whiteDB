package store

import (
	"os"
)

const (
	KB      = 1024
	MB      = 1024 * 1024
	LIMIT   = 3 * MB
	MAX_LEN = 2*LIMIT + MB // 3MB
)

type Set struct {
	Key   string
	Value string
}

type Remove struct {
	Key string
}

type Get struct {
	Key   string
	Value interface{}
}

type KVstore struct {
	ReadWriter *os.File
	Index      map[string]CommandPos
}

type CommandPos struct {
	Pos uint64
	Len uint64
}
