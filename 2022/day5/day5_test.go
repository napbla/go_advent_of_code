package main

import (
	"testing"
)

func TestParseString(t *testing.T) {
	a, b, c := parseMove("move 11 from 7 to 9")
	if a != 11 || b != 7 || c != 9 {
		t.Errorf("wrong parse has %d-%d-%d instead of 11,7,9", a, b, c)
	}
}
