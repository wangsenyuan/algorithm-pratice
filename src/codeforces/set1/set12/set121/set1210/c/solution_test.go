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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
4 5 6 0 8
1 2
1 3
1 4
4 5
`
	expect := 42
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
0 2 3 0 0 0 0
1 2
1 3
2 4
2 5
3 6
3 7
`
	expect := 30
	runSample(t, s, expect)
}
