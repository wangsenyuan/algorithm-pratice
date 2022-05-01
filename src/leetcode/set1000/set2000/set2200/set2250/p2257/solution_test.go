package p2257

import "testing"

func runSample(t *testing.T, m int, n int, guards [][]int, walls [][]int, expect int) {
	res := countUnguarded(m, n, guards, walls)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	m, n := 4, 6
	guards := [][]int{{0, 0}, {1, 1}, {2, 3}}
	walls := [][]int{{0, 1}, {2, 2}, {1, 4}}
	expect := 7
	runSample(t, m, n, guards, walls, expect)
}

func TestSample2(t *testing.T) {
	m, n := 3, 3
	guards := [][]int{{1, 1}}
	walls := [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}
	expect := 4
	runSample(t, m, n, guards, walls, expect)
}
