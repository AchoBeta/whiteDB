package store

import (
	"NekoKV/pkg/warn"
	"os"
	"sync"
)

var lock sync.Mutex

func Compact() {
	rw := Kvstore.ReadWriter
	ri := Kvstore.Index
	lock.Lock()
	defer lock.Unlock()
	// 切换到合并文件
	kv := Kvstore
	kv.CurPage = getRWFilePage(kv.CurPage)
	kv.SwitchFile()

	for k, row := range ri {
		data := read(rw, int64(row.Pos), int(row.Len))
		if data == nil {
			continue
		}
		pos, err := kv.Seek()
		if err != nil {
			warn.ERRORF(err.Error())
			return
		}
		kv.WriterAt(pos, string(data))
		// 记录新的索引
		newpos, _ := kv.Seek()
		kv.Index[k] = CommandPos{
			Pos:  uint64(pos),
			Len:  uint64(newpos - pos),
			Page: uint(kv.CurPage),
		}
	}
	err := os.Truncate(rw.Name(), 0) // 清空
	if err != nil {
		warn.ERRORF(err.Error())
		return
	}
}

func read(rw *os.File, pos int64, len int) []byte {
	data := make([]byte, len)
	n, err := rw.ReadAt(data, pos)
	if err != nil {
		warn.ERRORF(err.Error())
		return nil
	}
	return data[:n]
}

func getRWFilePage(page int) int {
	// fmt.Printf("page:%d, page ^= %d", page, page^1)
	return page ^ 1
}
