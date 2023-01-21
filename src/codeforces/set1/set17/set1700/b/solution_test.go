package main

import (
	"bytes"
	"testing"
)

func runSample(t *testing.T, s string) {
	res := solve(s)

	var buf bytes.Buffer

	n := len(s)
	var carray int
	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - '0')
		y := int(res[i] - '0')
		z := (x + y + carray) % 10
		carray = (x + y + carray) / 10
		buf.WriteByte(byte('0' + z))
	}
	if carray > 0 {
		buf.WriteByte(byte('0' + carray))
	}

	for i, j := 0, buf.Len()-1; i < j; i, j = i+1, j-1 {
		if buf.Bytes()[i] != buf.Bytes()[j] {
			t.Fatalf("Sample result %s, get non-palidrome addition %s", res, buf.String())
		}
	}
}

func TestSample1(t *testing.T) {
	s := "8189048033"
	runSample(t, s)
}
