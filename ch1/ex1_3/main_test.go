package main

import (
	"testing"
)

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsJoin()
	}
}

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringPlus()
	}
}
