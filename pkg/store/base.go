package store

import (
	"os"
)

/** 常量定义 **/

const (
	KB      = 1024
	MB      = 1024 * 1024
	LIMIT   = 3 * MB       // 3MB
	MAX_LEN = 2*LIMIT + MB // 单条命令+结构最大5M
)

const (
	dataFile    = "./datafile/DB_data.txt"
	compactFile = "./datafile/DB_compact.txt"
	tempFile    = "./datafile/DB_temp.txt"
)

const (
	DATA_FILE_PAGE    int = 0
	COMPACT_FILE_PAGE int = 1
	TEMP_FILE_PAGE    int = 2
)

/** 变量定义 **/

var fileCollection = []string{dataFile, compactFile, tempFile}

/** 结构体定义 **/

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
	CurPage    int
	CommandLen int
}

type CommandPos struct {
	Pos  uint64
	Len  uint64
	Page uint
}
