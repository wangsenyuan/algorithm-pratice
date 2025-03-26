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
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `10 8
1 6 2 3 2 6 3 10 1 2
1 10
2 2
3 3
2 3
1 3
3 6
4 6
5 5
`
	expect := `8
1
1
3
5
3
1
0
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1
1 3 2
1 3
`
	expect := `3`
	runSample(t, s, expect)
}
