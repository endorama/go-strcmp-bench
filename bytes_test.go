package main

import (
	"bytes"
	"testing"
)

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
		strcmp_bytes(strB, strA) // >
		strcmp_bytes(strA, strA) // ==
	}
}

func TestBytesWithBytes(t *testing.T) {
	if bytes.Compare(bytA, bytB) != -1 {
		t.Error("wrong -1")
	}
	if bytes.Compare(bytB, bytA) != +1 {
		t.Error("wrong +1")
	}
	if bytes.Compare(bytA, bytA) != 0 {
		t.Error("wrong 0")
	}
}

func BenchmarkBytesWithBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes.Compare(bytA, bytB) // <
		bytes.Compare(bytB, bytA) // >
		bytes.Compare(bytA, bytA) // ==
	}
}
