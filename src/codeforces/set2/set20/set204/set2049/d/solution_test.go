package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3 100
3 4 9
5 2 4
0 101 101`
	expect := 113
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 4 1
10 0 0 10
0 0 10 0
10 10 0 10`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1 3
4`
	expect := 4
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 2 3
1 2
3 6
5 4`
	expect := 13
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10 10 14
58 49 25 12 89 69 8 49 71 23
45 27 65 59 36 100 73 23 5 84
82 91 54 92 53 15 43 46 11 65
61 69 71 87 67 72 51 42 55 80
1 64 8 54 61 70 47 100 84 50
86 93 43 51 47 35 56 20 33 61
100 59 5 68 15 55 69 8 8 60
33 61 20 79 69 51 23 24 56 28
67 76 3 69 58 79 75 10 65 63
6 64 73 79 17 62 55 53 61 58`
	expect := 618
	runSample(t, s, expect)
}