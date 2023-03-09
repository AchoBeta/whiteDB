package store

import (
	"os"
	"whiteDB/pkg/warn"
)

const fileName = "./datafile/DB.txt"

var Kvstore *KVstore

func init() {
	isExist()
	rw, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0754)
	if err != nil {
		warn.ERRORF(err.Error())
		return
	}
	Kvstore = &KVstore{
		ReadWriter: rw,
		Index:      make(map[string]CommandPos),
	}
}

func isExist() {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		err := os.MkdirAll("./datafile", 0755)
		if err != nil {
			warn.ERRORF(err.Error())
			return
		}
		_, err = os.Create(fileName)
		if err != nil {
			warn.ERRORF(err.Error())
			return
		}
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
