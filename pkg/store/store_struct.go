package store

import (
	"os"
)

type Set struct {
	Key   string
	Value string
}

type Remove struct {
	Key string
}

type KVstore struct {
	ReadWriter *os.File
	Index      map[string]CommandPos
}

type CommandPos struct {
	Pos uint64
	Len uint64
}
