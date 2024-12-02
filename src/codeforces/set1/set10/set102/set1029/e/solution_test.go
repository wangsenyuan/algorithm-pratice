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
	s := `7
1 2
2 3
2 4
4 5
4 6
5 7
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7
1 2
1 3
2 4
2 5
3 6
1 7
`
	expect := 0
	runSample(t, s, expect)
}
func TestSample3(t *testing.T) {
	s := `7
1 2
2 3
3 4
3 5
3 6
3 7
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `9
1 2
2 3
3 4
4 5
5 6
5 7
7 8
8 9
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `9
1 6
2 9
3 4
4 9
5 3
6 2
7 8
8 1
`
	expect := 2
	runSample(t, s, expect)
}
