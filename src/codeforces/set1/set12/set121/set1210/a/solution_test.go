package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4
1 2
2 3
3 4
4 1
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1
1 3
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 21
1 2
1 3
1 4
1 5
1 6
1 7
2 3
2 4
2 5
2 6
2 7
3 4
3 5
3 6
3 7
4 5
4 6
4 7
5 6
5 7
6 7
`
	expect := 16
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7 14
7 3
2 4
2 1
2 5
5 3
6 7
4 7
5 4
7 5
4 3
4 1
6 1
6 3
3 1
`
	expect := 13
	runSample(t, s, expect)
}
