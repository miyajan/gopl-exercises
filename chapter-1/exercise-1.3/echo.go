package echo

import (
	"fmt"
	"strings"
)

func EchoInefficiently(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func EchoWithStringsJoin(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}
