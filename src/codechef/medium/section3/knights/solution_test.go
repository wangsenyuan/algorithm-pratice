package main

import "testing"

func runSample(t *testing.T, m, n int, G []string, expect int) {
	g := make([][]byte, m)
	for i := 0; i < m; i++ {
		g[i] = []byte(G[i])
	}
	res := solve(m, n, g)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n := 2, 4
	g := []string{
		"....",
		"....",
	}
	expect := 4
	runSample(t, m, n, g, expect)
}

func TestSample2(t *testing.T) {
	m, n := 5, 5
	g := []string{
		"..#..",
		"#..#.",
		"##...",
		"...##",
		".....",
	}
	expect := 11
	runSample(t, m, n, g, expect)
}
