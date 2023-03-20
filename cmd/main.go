package main

import (
	"NekoKV/pkg/comd"
	"NekoKV/pkg/store"
	"flag"
)

func main() {
	flag.Parse()
	store.LoadData()
	comd.ExecComd()
}
