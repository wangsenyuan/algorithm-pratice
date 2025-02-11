package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
12 18 11
`
	expect := 16
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
324097321 555675086 304655177 991244276 9980291
`
	expect := 805306368
	runSample(t, s, expect)
}
