package p3281

import "testing"

func runSample(t *testing.T, kx int, ky int, positions [][]int, expect int) {
	res := maxMoves(kx, ky, positions)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	kx, ky := 1, 1
	positions := [][]int{{0, 0}}
	expect := 4
	runSample(t, kx, ky, positions, expect)
}

func TestSample2(t *testing.T) {
	kx, ky := 0, 2
	positions := [][]int{{1, 1}, {2, 2}, {3, 3}}
	expect := 8
	runSample(t, kx, ky, positions, expect)
}

func TestSample3(t *testing.T) {
	kx, ky := 0, 0
	positions := [][]int{{1, 2}, {2, 4}}
	expect := 3
	runSample(t, kx, ky, positions, expect)
}
