package main

import (
	"flag"
	"whiteDB/pkg/comd"
)

func main() {
	flag.Parse()
	comd.ExecComd()
}
