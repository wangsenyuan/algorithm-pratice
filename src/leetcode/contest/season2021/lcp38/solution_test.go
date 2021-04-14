package lcp38

import "testing"

func runSample(t *testing.T, grid []string, expect int) {
	res := guardCastle(grid)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{"S.C.P#P.", ".....#.S"}
	expect := 3
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{"SP#P..P#PC#.S", "..#P..P####.#"}
	expect := -1
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{"SP#.C.#PS", "P.#...#.P"}
	expect := 0
	runSample(t, grid, expect)
}

func TestSample4(t *testing.T) {
	grid := []string{"CP.#.P.", "...S..S"}
	expect := 4
	runSample(t, grid, expect)
}

func TestSample5(t *testing.T) {
	grid := []string{"..S..C", "S#...."}
	expect := 2
	runSample(t, grid, expect)
}
