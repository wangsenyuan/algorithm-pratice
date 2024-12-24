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
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2
2 3
3 4
3 5`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
1 2
2 3
3 4`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5
2 1
3 1
4 1
5 4`
	expect := 3
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6
2 1
3 1
4 1
5 3
6 3`
	expect := 4
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `6
2 1
3 2
4 2
5 3
6 4`
	expect := 3
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `20
2 1
3 2
4 3
5 4
6 5
7 6
8 7
9 6
10 8
11 10
12 11
13 11
14 12
15 13
16 13
17 11
18 17
19 17
20 17`
	expect := 7
	runSample(t, s, expect)
}
