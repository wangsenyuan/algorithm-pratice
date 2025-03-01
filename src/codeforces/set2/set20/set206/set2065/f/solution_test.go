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
	s := `3
1 2 3
1 3
2 3`
	expect := "000"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
3 1 1 3
1 2
2 3
4 2`
	expect := "1010"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
2 4 4 2
1 2
2 3
3 4`
	expect := "0001"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `13
1 4 4 7 4 7 1 1 7 11 11 11 11
1 2
2 3
3 4
4 5
4 6
2 7
7 8
2 9
6 10
5 11
11 12
10 13`
	expect := "1001001000100"
	runSample(t, s, expect)
}
