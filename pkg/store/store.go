package store

import (
	"os"
	"whiteDB/pkg/warn"
)

const fileName = "./file/DB.txt"

var Kvstore *KVstore

func init() {
	rw, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		warn.ERRORF(err.Error())
		return
	}
	Kvstore = &KVstore{
		ReadWriter: rw,
		Index:      make(map[string]CommandPos),
	}
}

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

func (kv *KVstore) ReadAt(offset int64) (string, int64, error) {
	data := make([]byte, 1024)
	n, err := kv.ReadWriter.ReadAt(data, offset)
	if err != nil {
		return "", 0, err
	}
	pos, err := kv.ReadWriter.Seek(0, 1)
	if err != nil {
		return "", 0, err
	}
	return string(data[:n]), pos, nil

}

func (kv *KVstore) Close() error {
	err := kv.ReadWriter.Close()
	if err != nil {
		return err
	}
	return nil
}
