package p865

import "testing"

func runSample(t *testing.T, grid []string, expect int) {
	res := shortestPathAllKeys(grid)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"@.a.#", "###.#", "b.A.B",
	}
	expect := 8
	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"@..aA", "..B#.", "....b",
	}
	expect := 6
	runSample(t, grid, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{
		"@...a", ".###A", "b.BCc",
	}
	expect := 10
	runSample(t, grid, expect)
}
