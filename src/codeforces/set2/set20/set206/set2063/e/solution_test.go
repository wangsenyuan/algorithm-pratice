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
	s := `3
1 2
1 3`
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 2
3 2`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
2 3
1 5
4 2
1 2`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `11
2 1
2 3
2 4
4 5
6 5
5 7
4 8
8 9
7 10
10 11`
	expect := 29
	runSample(t, s, expect)
}
