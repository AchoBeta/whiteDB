package comd

import (
	"NekoKV/pkg/run"
	"NekoKV/pkg/store"
	"NekoKV/pkg/warn"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golang/glog"
)

func ExecComd() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("NekoKV >> ")
	for input.Scan() {
		line := input.Text()
		parser(line)
		fmt.Printf("NekoKV >> ")
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
		if checkSet(str) {
			run.ExecSet(str[1], str[2])
		}
	case LEN:
		run.ExecLen()
	case REMOVE:
		run.ExecRemove(str[1])
	case GET:
		run.ExecGet(str[1])
	case COMPACT:
		store.Compact()
	case EXIT:
		warn.EXIT()
		os.Exit(0)
	default:
		warn.ERRORF(str[0] + " is error command !")
	}
}

func checkSet(exec []string) bool {
	if len(exec) < 3 {
		fmt.Printf("Set must hava a value!\n")
		return false
	}
	_, k, v := exec[0], []byte(exec[1]), []byte(exec[2])
	if len(k) > store.LIMIT {
		warn.ERRORF("This key is too big !!")
		return false
	}

	if len(v) > store.LIMIT {
		warn.ERRORF("This value is too big!!")
		return false
	}
	return true
}
