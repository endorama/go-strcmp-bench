package main

import "testing"

func TestBytes(t *testing.T) {
	if strcmp_bytes(strA, strB) != -1 {
		t.Error("wrong -1")
	}
	if strcmp_bytes(strB, strA) != +1 {
		t.Error("wrong +1")
	}
	if strcmp_bytes(strA, strA) != 0 {
		t.Error("wrong 0")
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcmp_bytes(strA, strB) // <
	}
	for i := 0; i < b.N; i++ {
		strcmp_bytes(strB, strA) // >
	}
	for i := 0; i < b.N; i++ {
		strcmp_bytes(strA, strA) // ==
	}
}
