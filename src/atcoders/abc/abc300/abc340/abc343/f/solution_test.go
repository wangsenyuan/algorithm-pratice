package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	reader = bufio.NewReader(strings.NewReader(expect))
	for _, x := range res {
		y := readNum(reader)
		if x != y {
			t.Fatalf("Sample expect %s, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5 4
3 3 1 4 5
2 1 3
2 5 5
1 3 3
2 2 4
`
	expect := `1
0
2
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 1
1000000000
2 1 1

`
	expect := `0`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 9
2 4 4 3 9 1 1 2
1 5 4
2 7 7
2 2 6
1 4 4
2 2 5
2 2 7
1 1 1
1 8 1
2 1 8
`
	expect := `0
1
0
2
4
`
	runSample(t, s, expect)
}
