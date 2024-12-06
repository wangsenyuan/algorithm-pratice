package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	count := func(edges [][]int) int {
		var res int
		for _, e := range edges {
			res += e[2]
		}
		return res
	}

	a := count(expect)
	b := count(res)

	if a != b {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 2
1 2 10
2 3 5`
	expect := [][]int{
		{1, 2, 5},
		{1, 3, 5},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 2 10
2 3 15
3 1 10
`
	expect := [][]int{
		{2, 3, 5},
	}
	runSample(t, s, expect)
}
