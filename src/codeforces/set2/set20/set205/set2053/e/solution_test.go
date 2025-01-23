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
	s := `2
1 2`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2
1 3
2 4
2 5`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `12
1 6
11 2
4 8
12 3
2 7
6 12
8 1
2 3
5 12
9 2
10 3`
	expect := 40
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10
1 2
2 3
3 4
4 5
5 6
4 7
6 8
4 9
4 10`
	expect := 27
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `25
1 16
11 22
6 14
3 1
20 14
23 17
25 19
10 11
3 18
10 6
2 21
4 5
11 12
4 9
9 13
8 6
6 1
3 7
8 19
10 24
15 13
1 2
3 4
17 8`
	expect := 171
	runSample(t, s, expect)
}
