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
	s := `3 1
2 3 1
1 2
2 3`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1
3 6 3
1 2
2 3`
	expect := 8
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 1
-2 -3 -1
1 2
2 3`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 1
5 -4 3 6 7 3
4 1
5 1
3 5
3 6
1 2`
	expect := 17
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `8 1
3 5 2 7 8 5 -3 -4
7 3
1 8
4 3
3 5
7 6
8 7
2 1`
	expect := 26
	runSample(t, s, expect)
}
