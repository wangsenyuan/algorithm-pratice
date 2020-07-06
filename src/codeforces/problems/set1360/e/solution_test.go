package main

import "testing"

func runSample(t *testing.T, n int, G []string, expect bool) {
	g := make([][]byte, n)
	for i := 0; i < n; i++ {
		g[i] = []byte(G[i])
	}

	res := solve(n, g)

	if res != expect {
		t.Errorf("Sample %d %v, expect %t, but got %t", n, G, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	G := []string{
		"0010",
		"0011",
		"0000",
		"0000",
	}
	expect := true
	runSample(t, n, G, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	G := []string{
		"10",
		"01",
	}
	expect := false
	runSample(t, n, G, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	G := []string{
		"00",
		"00",
	}
	expect := true
	runSample(t, n, G, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	G := []string{
		"0101",
		"1111",
		"0101",
		"0111",
	}
	expect := true
	runSample(t, n, G, expect)
}

func TestSample5(t *testing.T) {
	n := 4
	G := []string{
		"0100",
		"1110",
		"0101",
		"0111",
	}
	expect := false
	runSample(t, n, G, expect)
}
