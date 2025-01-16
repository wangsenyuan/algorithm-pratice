package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, x string, expect string) {
	_, adj, s, res, path := process(bufio.NewReader(strings.NewReader(x)))

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
	if expect != "Win" {
		return
	}
	if path[0] != s {
		t.Fatalf("Sample result %v, not correct", path)
	}

	u := s
	for i := 1; i < len(path); i++ {
		v := path[i]
		ok := false
		for _, w := range adj[u-1] {
			if v == w {
				ok = true
				break
			}
		}
		if !ok {
			t.Fatalf("Sample result %v, not correct", path)
		}
		u = v
	}
	if len(adj[u-1]) != 0 {
		t.Fatalf("Sample result %v, not correct", path)
	}
}

func TestSample1(t *testing.T) {
	s := `5 6
2 2 3
2 4 5
1 4
1 5
0
1
`
	expect := "Win"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2
1 3
1 1
0
2
`
	expect := "Lose"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2
1 2
1 1
1
`
	expect := "Draw"
	runSample(t, s, expect)
}

