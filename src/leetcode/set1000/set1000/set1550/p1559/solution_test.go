package p1559

import "testing"

func runSample(t *testing.T, grid []string, expect bool) {
	G := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++ {
		G[i] = []byte(grid[i])
	}
	res := containsCycle(G)

	if res != expect {
		t.Errorf("Sample %v, expect %t, but got %t", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"aaaa",
		"abba",
		"abba",
		"aaaa",
	}
	expect := true

	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"ccca",
		"cdcc",
		"ccec",
		"fccc",
	}
	expect := true

	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		"abb",
		"bzb",
		"bba",
	}
	expect := false

	runSample(t, grid, expect)
}
