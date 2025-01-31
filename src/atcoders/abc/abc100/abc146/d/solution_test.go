package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, k, res := process(reader)
	if k != expect {
		t.Fatalf("Sample %s, expect %d, but got %d", s, expect, k)
	}
	g := NewGraph(n, 2*n)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	for u := 0; u < n; u++ {
		cnt := make(map[int]int)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			j := g.val[i]
			cnt[res[j]]++
			if cnt[res[j]] > 1 {
				t.Fatalf("Sample result %v, not correct", cnt)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2
2 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8
1 2
2 3
2 4
2 5
4 7
5 6
6 8
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
1 2
1 3
1 4
1 5
1 6
`
	expect := 5
	runSample(t, s, expect)
}
