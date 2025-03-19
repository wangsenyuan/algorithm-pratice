package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2
1 1
6 1
1 6
1 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
4 4
4 5
5 4
4 4
`
	expect := false
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `3 3
1 2 3
4 5 6
7 8 9
1 4 7
2 5 6
3 8 9
`
	expect := true
	runSample(t, s, expect)
}
