package main

import (
	"strings"
	"testing"
)

func runSample(t *testing.T, n, m int, G []string, expect int) {
	g := make([][]byte, n)
	for i := 0; i < n; i++ {
		g[i] = []byte(G[i])
	}

	res := solve(n, m, g)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 3
	S := `.?.
.?.
.#.`
	G := strings.Split(S, "\n")
	expect := 8
	runSample(t, n, m, G, expect)
}

func TestSample2(t *testing.T) {
	n, m := 5, 8
	S := `.#...##.
.##..?..
.###.#.#
??#..?..
###?#...`
	G := strings.Split(S, "\n")
	expect := 42
	runSample(t, n, m, G, expect)
}
