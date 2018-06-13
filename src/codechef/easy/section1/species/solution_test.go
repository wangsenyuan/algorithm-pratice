package main

import "testing"

func runSample(t *testing.T, n int, grid []string, expect int) {
	gg := make([][]byte, n)
	for i := 0; i < n; i++ {
		gg[i] = []byte(grid[i])
	}
	res := solve(n, gg)

	if res != expect {
		t.Errorf("sample %d %v, expect %d, but got %d", n, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	grid := []string{
		"..?",
		".?B",
		"G..",
	}
	expect := 1
	runSample(t, n, grid, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	grid := []string{
		"GG",
		"..",
	}
	expect := 0
	runSample(t, n, grid, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	grid := []string{
		"?..",
		".??",
		"??.",
	}
	expect := 6
	runSample(t, n, grid, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	grid := []string{
		"??P",
		"???",
		"??B",
	}
	expect := 0
	runSample(t, n, grid, expect)
}

func TestSample5(t *testing.T) {
	n := 7
	grid := []string{
		"?.?.?.?",
		".?.?.?.",
		"?.?.?.?",
		".?.?.?.",
		"?.?.?.?",
		".?.?.?.",
		"?.?.?.?",
	}
	expect := 288603514
	runSample(t, n, grid, expect)
}

func TestSample6(t *testing.T) {
	n := 2
	grid := []string{
		"PP",
		"PP",
	}
	expect := 1
	runSample(t, n, grid, expect)
}

func TestSample7(t *testing.T) {
	n := 2
	grid := []string{
		"G.",
		"?.",
	}
	expect := 0
	runSample(t, n, grid, expect)
}

func TestSample8(t *testing.T) {
	n := 3
	grid := []string{
		"?.?",
		"?.?",
		"...",
	}
	expect := 4
	runSample(t, n, grid, expect)
}
func TestSample9(t *testing.T) {
	n := 3
	grid := []string{
		"?.?",
		"?.?",
		".?.",
	}
	expect := 12
	runSample(t, n, grid, expect)
}
