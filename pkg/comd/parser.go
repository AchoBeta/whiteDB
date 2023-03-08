package comd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"whiteDB/pkg/run"
	"whiteDB/pkg/warn"
)

const (
	NONE int = iota
	SET
	REMOVE
	GET
)

var commandMap map[string]int = map[string]int{
	"SET":    SET,
	"REMOVE": REMOVE,
	"GET":    GET,
}

func ExecComd() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("WhiteDB >> ")
	for input.Scan() {
		line := input.Text()
		parser(line)
		if line == "exit" {
			warn.EXIT()
			os.Exit(0)
		}
		fmt.Printf("%s \n", line)
		fmt.Printf("WhiteDB >> ")
	}
}

func parser(exec string) {
	str := strings.Fields(exec)
	cmd, ok := commandMap[str[0]]
	if !ok {
		cmd = NONE
	}
	switch cmd {
	case SET:
		run.ExecSet()
	case REMOVE:
	case GET:
	default:
		warn.ERROR()
	}
}
