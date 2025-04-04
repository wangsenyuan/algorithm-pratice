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
	s := `4 1 1
4
1 2
2 3
3 4
4 1
4
1 2
2 3
3 4
4 1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 1 2
4
1 2
2 3
3 4
4 1
4
1 2
2 3
3 4
4 1`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 7 2
7
1 6
2 1
3 2
3 4
5 1
7 3
7 5
6
5 1
5 6
5 7
6 3
7 2
7 4`
	expect := 7
	runSample(t, s, expect)
}
