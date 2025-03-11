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
		t.Fatalf("Sampole expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 2
2 0 0 1 2
1 2
2 3
2 4
2 5
`
	expect := 1
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := `7 3
0 1 0 2 2 3 0
1 3
1 4
1 5
2 7
3 6
4 7
`
	expect := 4
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `2 2
1 2
1 2
`
	expect := 1
	runSample(t, s, expect)
}


func TestSample4(t *testing.T) {
	s := `10 5
0 1 0 2 3 0 1 4 5 1
2 3
3 1
2 10
10 6
7 10
9 4
5 9
5 6
9 8
`
	expect := 2
	runSample(t, s, expect)
}
