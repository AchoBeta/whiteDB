package store

import (
	"NekoKV/pkg/warn"
	"bufio"
	"encoding/json"
)

var (
	set    Set
	remove Remove
)

func LoadData() {
	kv := Kvstore
	sc := bufio.NewScanner(kv.ReadWriter)
	sc.Split(bufio.ScanLines)
	pos := 0
	stat, _ := (kv.ReadWriter).Stat()
	if stat.Size() == 0 {
		kv.SwapCurPage()
		kv.SwitchFile()
		sc = bufio.NewScanner(kv.ReadWriter)
		sc.Split(bufio.ScanLines)
	}
	for sc.Scan() {
		debug(sc)
		data := sc.Bytes()
		len := len(data)

		load(kv, []byte(data), uint64(pos), uint64(len))
		pos += len + 1
	}
}

func load(kv *KVstore, buf []byte, pos uint64, len uint64) {

	var m map[string]interface{}
	err := json.Unmarshal(buf, &m)
	if err != nil {
		warn.ERRORF(err.Error())
		return
	}
	// 根据map中是否有Value字段判断是Set还是Remove结构体
	if _, ok := m["Value"]; ok {
		// 如果有Value字段，则转换成Set结构体
		if err := json.Unmarshal(buf, &set); err == nil {
			kv.Index[set.Key] = CommandPos{
				Pos: uint64(pos),
				Len: len,
			}
			return
		}
	} else {
		if err := json.Unmarshal(buf, &remove); err == nil {
			delete(kv.Index, remove.Key)
			return
		}
	}

}

func debug(sc *bufio.Scanner) {
	// line := sc.Text()
	// fmt.Printf(line+"=> pos:%d, len:%d\n", pos, len)
}
