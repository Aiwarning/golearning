package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg + arg
		sep = " "
	}
	fmt.Println(s)
}
