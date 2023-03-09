package main

import (
	"flag"
	"whiteDB/pkg/comd"
	"whiteDB/pkg/store"
)

func main() {
	flag.Parse()
	store.LoadData()
	comd.ExecComd()
}
