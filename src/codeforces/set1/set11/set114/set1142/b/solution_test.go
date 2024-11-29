package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 6 3
2 1 3
1 2 3 1 2 3
1 5
2 6
3 5
`
	expect := "110"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 4 3
2 1
1 1 2 2
1 2
2 3
3 4
`
	expect := "010"
	runSample(t, s, expect)
}
