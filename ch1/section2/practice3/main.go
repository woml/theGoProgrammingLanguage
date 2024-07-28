package main

import (
	"strings"
)

func main() {

}

func stringPlus(inputs []string) {
	var s, sep string
	for _, input := range inputs {
		s += sep + input
		sep = " "
	}
}

func stringJoin(inputs []string) {
	strings.Join(inputs, " ")
}
