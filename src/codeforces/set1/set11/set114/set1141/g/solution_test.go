package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	n, k, edges, r, res := process(reader)

	if r != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, r)
	}
	g := make([]map[int]int, n)
	for i := range n {
		g[i] = make(map[int]int)
	}
	var cnt int
	deg := make([]int, n)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u][res[i]]++
		g[v][res[i]]++
		deg[u]++
		deg[v]++
	}

	for u := 0; u < n; u++ {
		if len(g[u]) != deg[u] {
			cnt++
		}
	}

	if cnt > k {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 2
1 4
4 3
3 5
3 6
5 2
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
3 1
1 4
1 2
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 2
10 3
1 2
1 3
1 4
2 5
2 6
2 7
3 8
3 9
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 0
2 3
2 4
3 1
`
	expect := 2
	runSample(t, s, expect)
}
