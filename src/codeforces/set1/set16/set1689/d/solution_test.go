package main

import "testing"

func runSample(t *testing.T, n int, m int, G []string, expect []int) {
	res := solve(n, m, G)

	if get_distance(G, expect) != get_distance(G, res) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func get_distance(G []string, cell []int) int {
	var res int

	for i := 0; i < len(G); i++ {
		for j := 0; j < len(G[i]); j++ {
			if G[i][j] == 'B' {
				res = max(res, abs(cell[0]-1-i)+abs(cell[1]-1-j))
			}
		}
	}

	return res
}

func TestSample1(t *testing.T) {
	n, m := 3, 2
	G := []string{
		"BW",
		"WW",
		"WB",
	}
	expect := []int{2, 1}

	runSample(t, n, m, G, expect)
}

func TestSample2(t *testing.T) {
	n, m := 3, 3
	G := []string{
		"WWB",
		"WBW",
		"BWW",
	}
	expect := []int{2, 2}

	runSample(t, n, m, G, expect)
}

func TestSample3(t *testing.T) {
	n, m := 5, 5
	G := []string{
		"BWBWB",
		"WBWBW",
		"BWBWB",
		"WBWBW",
		"BWBWB",
	}
	expect := []int{3, 3}

	runSample(t, n, m, G, expect)
}

func TestSample4(t *testing.T) {
	n, m := 9, 9
	G := []string{
		"WWWWWWWWW",
		"WWWWWWWWW",
		"BWWWWWWWW",
		"WWWWWWWWW",
		"WWWWBWWWW",
		"WWWWWWWWW",
		"WWWWWWWWW",
		"WWWWWWWWW",
		"WWWWWWWWB",
	}
	expect := []int{6, 5}

	runSample(t, n, m, G, expect)
}
