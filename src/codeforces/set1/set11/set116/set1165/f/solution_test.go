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
	s := `5 6
1 2 0 2 0
2 4
3 3
1 5
1 2
1 5
2 3
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3
4 2 1 3 2
3 5
4 2
2 5
`
	expect := 20
	runSample(t, s, expect)
}
