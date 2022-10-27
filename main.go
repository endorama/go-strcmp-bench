package main

import (
	"bytes"
	"strings"
)

const strA = "foobar"
const strB = "hello"

var bytA = []byte(strA)
var bytB = []byte(strB)

func strcmp_builtin(a, b string) int {
	if a > b {
		return +1
	}
	if a < b {
		return -1
	}
	// a == b
	return 0
}

func strcmp_strings(a, b string) int {
	return strings.Compare(a, b)
}

func strcmp_bytes(a, b string) int {
	c := []byte(a)
	d := []byte(b)

	return bytes.Compare(c, d)
}
