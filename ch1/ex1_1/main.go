// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// gopl.io/ch1/echo1 を改変したものです。

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
