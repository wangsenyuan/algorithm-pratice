package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1263 1 2
8103 2 1
`
	expect := "Need more data"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1234 2 2
1256 0 2
`
	expect := "2134"
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `2
0123 1 1
4567 1 2
`
	expect := "Incorrect data"
	runSample(t, s, expect)
}
