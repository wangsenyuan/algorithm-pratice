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
	s := `4
0 11
6 6
0 0
5 5
2 1
3 1
4 3`
	expect := 11
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
1 1
0 5
0 5
2 2
2 2
2 2
2 2
1 2
1 3
2 4
2 5
3 6
3 7`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7
0 20
0 20
0 20
0 20
3 3
4 4
5 5
1 2
1 3
1 4
2 5
3 6
4 7`
	expect := 5
	runSample(t, s, expect)
}
