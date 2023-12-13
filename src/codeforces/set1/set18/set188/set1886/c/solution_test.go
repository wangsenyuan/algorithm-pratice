package main

import "testing"

func runSample(t *testing.T, s string, pos int, expect byte) {
	res := solve(s, pos)

	if res != expect {
		t.Fatalf("Sample expect %c, but got %c", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "cab"
	pos := 6
	var expect byte = 'a'
	runSample(t, s, pos, expect)
}

func TestSample2(t *testing.T) {
	s := "abcd"
	pos := 9
	var expect byte = 'b'
	runSample(t, s, pos, expect)
}
