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
codeforces
codehorses
code
`
	expect := 6
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := `5
abba
abb
ab
aa
aacada
`
	expect := 11
	runSample(t, s, expect)
}
