package run

import (
	"encoding/json"
	"fmt"
	"sync"

	"whiteDB/pkg/store"
	"whiteDB/pkg/warn"
)

var (
	lock sync.Mutex
)

func ExecSet(key, value string) {
	set := &store.Set{
		Key:   key,
		Value: value,
	}
	kv := store.Kvstore
	// lock
	lock.Lock()
	defer lock.Unlock()
	pos, err := kv.Seek()
	if err != nil {
		return
	}

	data, err := json.Marshal(set)
	data = append(data, '\n')
	if err != nil {
		warn.ERRORF(err.Error())
		return
	}
	kv.WriterAt(pos, string(data))
	// 记录索引
	npos, _ := kv.Seek()
	kv.Index[key] = store.CommandPos{
		Pos:  uint64(pos),
		Len:  uint64(npos - pos),
		Page: uint(kv.CurPage),
	}
}

func ExecRemove(key string) {
	kv := store.Kvstore
	rm := &store.Remove{
		Key: key,
	}
	data, err := json.Marshal(rm)
	data = append(data, '\n')
	if err != nil {
		warn.ERRORF(err.Error())
		return
	}
	// lock
	lock.Lock()
	defer lock.Unlock()
	kv.Writer(string(data))
	// 索引中删除
	delete(kv.Index, key)
}

func ExecGet(key string) {
	kv := store.Kvstore
	var val string = "nil"
	// 从索引中取出数据
	if cmd, ok := kv.Index[key]; ok {
		pos, len := cmd.Pos, cmd.Len
		data, err := kv.ReadAt(int64(pos), int(len))
		if err != nil {
			warn.ERRORF(err.Error())
			return
		}
		get := &store.Get{}
		err = json.Unmarshal(data, get)
		if err != nil {
			warn.ERRORF(err.Error())
			return
		}
		val = get.Value.(string)
	}
	fmt.Printf(val + "\n")
}
