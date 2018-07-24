package main

import "testing"

func runSample(t *testing.T, R, C int, G []string, expect bool) {
	g := make([][]byte, R)
	for i := 0; i < R; i++ {
		g[i] = []byte(G[i])
	}
	res := solve(R, C, g)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %t, but got %t", R, C, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	R, C := 2, 2
	G := []string{
		"K.", ".#",
	}
	runSample(t, R, C, G, false)
}

func TestSample2(t *testing.T) {
	R, C := 4, 2
	G := []string{
		"K#", ".#", ".#", ".#",
	}
	runSample(t, R, C, G, true)
}

func TestSample3(t *testing.T) {
	R, C := 4, 2
	G := []string{
		"##", ".#", "..", "K#",
	}
	runSample(t, R, C, G, true)
}
