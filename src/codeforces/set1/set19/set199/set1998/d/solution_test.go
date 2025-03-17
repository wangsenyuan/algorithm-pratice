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
	s := `6 0
`
	expect := "11111"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 1
2 6
`
	expect := "11011"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 1
1 5
`
	expect := "10011"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 4
1 3
1 6
2 7
3 8
`
	expect := "100001111"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `15 3
2 8
4 9
8 15
`
	expect := "11000111000111"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `27 13
17 27
13 15
21 25
4 17
25 27
8 25
19 24
15 25
5 13
14 24
12 15
3 5
7 19
`
	expect := "11110000000000000000011111"
	runSample(t, s, expect)
}
