package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i] // os.Args[0]にはコマンド自身が出力される
		sep = " "
	}
	fmt.Println(s)
}
