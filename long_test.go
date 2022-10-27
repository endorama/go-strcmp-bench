package main

import (
	"bytes"
	"testing"
)

const longStrA = "QTlTUW5NyjpDubyeGNRK5hMyKdXdYIboSRBr4jQOW5TyxLNWrf1VFxOR8PyhBpAVnJRMfqAray3fnTQDBIp3wzNB0ik1ga0HTCVt6hnKUKqIZfBCi0MqkkUtOCpklN1JUktpXb7LHaM72nVYbEbvTlLCRC5Sf2jNN9f4q5v1pdnKdrWEYjpUc48bw4uTK6Ero6TA8pcBn1vgRoYWZfhHZ2QtBGSzfQqX34ZZmGB6FVtICqVbhdElXZxtaaDGM0q"
const longStrB = "8GEEllZiGs639zhbu2zvDTlBlnWlJC7f2tV4onGNonYgTbrngPYawZ4El9dD7anjNhRTQ4g8f6q0XfG1ukzY8JChVKmj6ehvkYXkGmVIcqUboTAuUi6d1WV7j5T7DOsUYsIL2YA4u8H9cO2h60MwcsG8wW0Pg04JS8XnwloPRKfvLoBriVfWKrop9srrmY7oS2noQjr2XI5tQG56SoQZOJEQAGcTDlpBwujdl5nbpMLLcHYyUWkAKDydtDh5WkP"

var longBytA = []byte(longStrA)
var longBytB = []byte(longStrB)

func BenchmarkBuiltIn_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcmp_builtin(longStrA, longStrB) // <
		strcmp_builtin(longStrB, longStrA) // >
		strcmp_builtin(longStrA, longStrA) // ==
	}
}

func BenchmarkBytes_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcmp_bytes(longStrA, longStrB) // <
		strcmp_bytes(longStrB, longStrA) // >
		strcmp_bytes(longStrA, longStrA) // ==
	}
}

func BenchmarkBytesWithBytes_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes.Compare(longBytA, longBytB) // <
		bytes.Compare(longBytB, longBytA) // >
		bytes.Compare(longBytA, longBytA) // ==
	}
}

func BenchmarkStrings_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strcmp_strings(longStrA, longStrB) // <
		strcmp_strings(longStrB, longStrA) // >
		strcmp_strings(longStrA, longStrA) // ==
	}
}
