package comd

import (
	"bufio"
	"fmt"
	"os"
)

func ExecComd() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("WhiteDB >> ")
	for input.Scan() {
		line := input.Text()
		if line == "exit" {
			os.Exit(0)
		}
		fmt.Printf("%s \n", line)
		fmt.Printf("WhiteDB >> ")
	}
}
