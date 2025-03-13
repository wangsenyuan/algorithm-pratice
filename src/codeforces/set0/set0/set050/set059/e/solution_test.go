package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res, n, roads, forbiden := process(reader)

	if len(res)-1 != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
	if expect < 0 {
		return
	}
	// let verify it
	if res[0] != 1 || res[expect] != n {
		t.Fatalf("Sample result %v, not correct", res)
	}

	bad := make(map[record]bool)
	for _, cur := range forbiden {
		r := record{cur[0] - 1, cur[1] - 1, cur[2] - 1}
		bad[r] = true
	}
	g := make([]map[int]bool, n+1)
	for i := range n + 1 {
		g[i] = make(map[int]bool)
	}

	for _, road := range roads {
		u, v := road[0], road[1]
		g[u][v] = true
		g[v][u] = true
	}

	cur := 1

	for i := 1; i < len(res); i++ {
		v := res[i]
		if !g[cur][v] {
			t.Fatalf("Sample result %v, not correct, it has no raod from %d => %d", res, cur, v)
		}
		if i >= 2 && bad[record{res[i-2], res[i-1], res[i]}] {
			t.Fatalf("Sample result %v, not correct, it passed a forbiden triple", res)
		}
		cur = v
	}
}

func TestSample1(t *testing.T) {
	s := `4 4 1
1 2
2 3
3 4
1 3
1 4 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 1 0
1 2
`
	expect := -1
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `4 4 2
1 2
2 3
3 4
1 3
1 2 3
1 3 4
`
	expect := 4
	runSample(t, s, expect)
}
