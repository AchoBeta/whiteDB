package run

import (
	"whiteDB/pkg/store"
)

func ExecSet(key, value string) {
	set := &store.Set{
		Key:   key,
		Value: value,
	}
	wr := store.Kvstore.Writer
}
