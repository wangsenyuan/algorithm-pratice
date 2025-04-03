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
01110
10010
10001
10011
11110
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10
0000000000
0000110000
1001000000
1000011110
1011111101
1011110011
1011000111
1011000001
1111000010
0000111110
`
	expect := 20
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
10
11
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
001
001
000
`
	expect := 4
	runSample(t, s, expect)
}

