// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// gopl.io/ch1/echo1 を改変したものです。
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", " "
	// rangeを使う場合、配列（スライス）の添字と値を変数に格納することが可能。
	for k, arg := range os.Args[1:] {
		s = strconv.Itoa(k)
		s += sep + arg
		fmt.Println(s)
	}
}
