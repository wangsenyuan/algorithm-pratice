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

func TestSamaple1(t *testing.T) {
	s := `3
1 2 1
1 2
2 3
`
	expect := `5
4
0
`
	runSample(t, s, expect)
}

func TestSamaple2(t *testing.T) {
	s := `1
1
`
	expect := `1`
	runSample(t, s, expect)
}

func TestSamaple3(t *testing.T) {
	s := `2
1 2
1 2
`
	expect := `2
2`
	runSample(t, s, expect)
}

func TestSamaple4(t *testing.T) {
	s := `5
1 2 3 4 5
1 2
2 3
3 4
3 5
`
	expect := `5
8
10
5
5
`
	runSample(t, s, expect)
}

func TestSamaple5(t *testing.T) {
	s := `8
2 7 2 5 4 1 7 5
3 1
1 2
2 7
4 5
5 6
6 8
7 8
`
	expect := `18
15
0
14
23
0
23
0
`
	runSample(t, s, expect)
}
