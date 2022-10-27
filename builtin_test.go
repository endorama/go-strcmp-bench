package main

import "testing"

func TestFirst(t *testing.T) {
	if strcmp_builtin(strA, strB) != -1 {
		t.Error("wrong -1")
	}
	if strcmp_builtin(strB, strA) != +1 {
		t.Error("wrong +1")
	}
	if strcmp_builtin(strA, strA) != 0 {
		t.Error("wrong 0")
	}
}

func BenchmarkBuiltIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcmp_builtin(strA, strB) // <
	}
	for i := 0; i < b.N; i++ {
		strcmp_builtin(strB, strA) // >
	}
	for i := 0; i < b.N; i++ {
		strcmp_builtin(strA, strA) // ==
	}
}
