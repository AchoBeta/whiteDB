package store

import (
	"NekoKV/pkg/warn"
	"os"
)

var (
	Kvstore *KVstore
)

func init() {
	isExist()
	rw := open(dataFile)
	if rw == nil {
		return
	}
	Kvstore = &KVstore{
		ReadWriter: rw,
		Index:      make(map[string]CommandPos),
		CurPage:    0,
	}
}

// open 返回指定文件的操作指针
func open(file string) *os.File {
	rw, err := os.OpenFile(file, os.O_APPEND|os.O_RDWR, 0754)
	if err != nil {
		warn.ERRORF(err.Error())
		return nil
	}
	return rw
}

// isExist 检查文件是否存在
func isExist() {

	for _, file := range fileCollection {
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			err := os.MkdirAll("./datafile", 0754)
			if err != nil {
				warn.ERRORF(err.Error())
				return
			}
			_, err = os.Create(file)
			if err != nil {
				warn.ERRORF(err.Error())
				return
			}
		}
	}
}

// SwitchFile 切换操作的文件
func (kv *KVstore) SwitchFile() {
	rw := open(fileCollection[kv.CurPage])
	if rw == nil {
		return
	}
	kv.ReadWriter = rw
}

// Seek 获取当前读取的位置
func (kv *KVstore) Seek() (int64, error) {
	pos, err := kv.ReadWriter.Seek(0, 1)
	if err != nil {
		warn.ERRORF(err.Error())
		return 0, err
	}
	return pos, nil
}

func (kv *KVstore) WriterAt(pos int64, data string) error {
	_, err := kv.ReadWriter.Seek(pos, 0)
	if err != nil {
		warn.ERRORF(err.Error())
		return err
	}
	_, err = kv.ReadWriter.WriteString(data)
	if err != nil {
		warn.ERRORF(err.Error())
		return err
	}
	return nil
}

func (kv *KVstore) Writer(data string) (int64, error) {
	_, err := kv.ReadWriter.WriteString(data)
	if err != nil {
		warn.ERRORF(err.Error())
		return 0, err
	}
	pos, err := kv.ReadWriter.Seek(0, 1) // 获取偏移量
	if err != nil {
		warn.ERRORF(err.Error())
		return 0, err
	}
	return pos, nil
}

// ReadAt 从offset开始读取len长度的内容
func (kv *KVstore) ReadAt(offset int64, len int) ([]byte, error) {
	data := make([]byte, len)
	n, err := kv.ReadWriter.ReadAt(data, offset)
	if err != nil {
		return nil, err
	}
	return data[:n], nil

}

func (kv *KVstore) Close() error {
	err := kv.ReadWriter.Close()
	if err != nil {
		return err
	}
	return nil
}

func (kv *KVstore) SwapCurPage() error {
	kv.CurPage = kv.CurPage ^ 1
	return nil
}
