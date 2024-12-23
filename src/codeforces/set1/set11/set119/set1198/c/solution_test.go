package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	n, edges, kind, res := process(reader)

	if !expect {
		if kind != imp {
			t.Fatalf("Sample expect %t, but got %s %v", expect, kind, res)
		}
		return
	}

	if kind == imp {
		t.Fatalf("Sample expect %t, but got %s %v", expect, kind, res)
	}
	// not imp
	if kind == mtc {
		ok := make([]bool, len(edges))
		for _, i := range res {
			ok[i-1] = true
		}
		marked := make([]bool, 3*n+1)

		for i, e := range edges {
			if ok[i] {
				u, v := e[0], e[1]
				if marked[u] || marked[v] {
					t.Fatalf("Sample result %v, not correct", res)
				}
				marked[u] = true
				marked[v] = true
			}
		}
	} else {
		// ind
		marked := make([]bool, 3*n+1)
		for _, i := range res {
			marked[i] = true
		}
		for _, e := range edges {
			u, v := e[0], e[1]
			if marked[u] && marked[v] {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `1 2
1 3
1 2`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 5
1 2
3 1
1 4
5 1
1 6`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 15
1 2
1 3
1 4
1 5
1 6
2 3
2 4
2 5
2 6
3 4
3 5
3 6
4 5
4 6
5 6`
	expect := true
	runSample(t, s, expect)
}
