package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(stringsJoin())
}

func stringPlus() string {
	// forループ、+=演算子を使った場合
	// ベンチマーク　14495076485 ns
	var str []string
	for i := 0; i < 100000; i++ {
		str = append(str, "hogehoge")
	}
	var s, sep string
	sep = " "
	for i := 0; i < len(str); i++ {
		// +=演算子を使って文字列結合
		s += sep + str[i]
	}
	return s
}

func stringsJoin() string {
	// strings.Joinを使った場合
	//ベンチマーク　405192894 ns
	var str []string
	for i := 0; i < 100000; i++ {
		str = append(str, "hogehoge")
	}
	// strings.Joinを使って文字列結合
	return strings.Join(str, " ")
}
