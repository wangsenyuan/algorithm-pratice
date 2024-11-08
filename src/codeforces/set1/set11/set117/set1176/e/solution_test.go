package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, n, edges := process(reader)

	if len(res)*2 > n {
		t.Fatalf("Sample result %v, picked too many nodes", res)
	}

	marked := make([]bool, n+1)
	ok := make([]bool, n+1)

	for _, x := range res {
		marked[x] = true
		ok[x] = true
	}

	for _, e := range edges {
		u, v := e[0], e[1]
		if marked[u] {
			ok[v] = true
		}
		if marked[v] {
			ok[u] = true
		}
	}

	for i := 1; i <= n; i++ {
		if !ok[i] {
			t.Fatalf("Sample result %v, not correct, as %d not covered", res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 6
1 2
1 3
1 4
2 3
2 4
3 4
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `6 8
2 5
5 4
4 3
4 1
1 3
2 3
2 6
5 6
`
	runSample(t, s)
}