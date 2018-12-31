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
