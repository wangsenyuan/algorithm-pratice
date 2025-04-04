package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 10
3 4
3 2
`
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 6
1 4
2 3
2 7
`
	expect := 14
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 10
1 7
`
	expect := 12
	runSample(t, s, expect)
}
