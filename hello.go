/* hello world and hello user */

package main

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	fmt.Println("Before main\n")
}

func main() {
	who := "World\n"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello ", who)
}
