package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res := process(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 5
3 1 2 3
3 3 4 5
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 10
1 1
1 2
1 3
1 4
1 5
1 6
1 7
1 8
1 9
1 10
`
	expect := false
	runSample(t, s, expect)
}
