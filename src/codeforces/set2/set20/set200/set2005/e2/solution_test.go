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
	s := `2 2 3
1 2
1 3 6
4 6 2`
	expect := "N"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2 4
1 2
1 1 3 2
4 2 5 1`
	expect := "T"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 4 2
1 2
3 4
5 6
7 8
8 8`
	expect := "N"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 6 6
1 2 3 4 5 6
1 1 7 7 7 7
7 2 7 7 7 7
7 7 3 7 7 7
7 7 7 4 7 7
7 7 7 7 5 7
7 7 7 7 7 6`
	expect := "T"
	runSample(t, s, expect)
}
