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
	s := `5 7
2 2 8 6 9
1 2
1 3
2 4
3 2
4 5
2 5
1 5
2
`
	expect := 27
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 12
1 7 1 9 3 3 6 30 1 10
1 2
1 3
3 5
5 7
2 3
5 4
6 9
4 6
3 7
6 8
9 4
9 10
6
`
	expect := 61
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 9
96 86 63 95 78 91 96 100 99 90
10 5
1 2
8 7
4 5
4 6
3 8
6 7
3 9
10 2
8
`
	expect := 732
	runSample(t, s, expect)
}
