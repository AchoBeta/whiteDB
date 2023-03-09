package comd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"whiteDB/pkg/run"
	"whiteDB/pkg/warn"

	"github.com/golang/glog"
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
	"RM":     REMOVE,
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
		fmt.Printf("WhiteDB >> ")
		glog.Flush()
	}
}

func parser(exec string) {
	str := strings.Fields(exec)
	cmd, ok := commandMap[strings.ToUpper(str[0])]
	if !ok {
		cmd = NONE
	}
	switch cmd {
	case SET:
		run.ExecSet(str[1], str[2])
	case REMOVE:
		run.ExecRemove(str[1])
	case GET:
	default:
		warn.ERRORF(str[0] + " is error command !")
	}
}
