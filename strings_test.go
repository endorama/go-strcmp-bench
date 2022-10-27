package main

import "testing"

func TestStrings(t *testing.T) {
	if strcmp_strings(strA, strB) != -1 {
		t.Error("wrong -1")
	}
	if strcmp_strings(strB, strA) != +1 {
		t.Error("wrong +1")
	}
	if strcmp_strings(strA, strA) != 0 {
		t.Error("wrong 0")
	}
}

func BenchmarkStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcmp_strings(strA, strB) // <
	}
	for i := 0; i < b.N; i++ {
		strcmp_strings(strB, strA) // >
	}
	for i := 0; i < b.N; i++ {
		strcmp_strings(strA, strA) // ==
	}
}
