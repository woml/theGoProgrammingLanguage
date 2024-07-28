package main

import (
	"testing"
	"theGoProgrammingLanguage"
)

var inputs = []string{"ab", "cd", "ef", "gh", "ij", "kl", "mn", "op", "qs", "yz", "uv", "wx", "z"}

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main.stringPlus(inputs)
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main.stringJoin(inputs)
	}
}
