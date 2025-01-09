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
	s := `5
1 2 3 4 5
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
8 12 7
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
3 2 1
`
	// 1  3  (3, 1)
	// 1, 2, 2 1

	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8
7 2 6 6 5 1 4 9
`
	// 1 (1, 3, 4, 5, 6, 7)
	// 2 (1, 2)
	// 3 (1, 2, 3, 4, 5, 6)
	// 4 (1, 2, 3, 4, 5, 6)
	// 5 (1, 2, 3, 4, 5)
	// 6 (1)
	// 7 (1, 2, 3, 4)
	// 8 (1, 2, 3, 4, 5, 6, 7, 8, 9)

	expect := 9
	runSample(t, s, expect)
}
