package p2056

import "testing"

func runSample(t *testing.T, pieces []string, positions [][]int, expect int) {
	res := countCombinations(pieces, positions)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pieces := []string{"rook"}
	positions := [][]int{{1, 1}}
	expect := 15
	runSample(t, pieces, positions, expect)
}

func TestSample2(t *testing.T) {
	pieces := []string{"queen"}
	positions := [][]int{{1, 1}}
	expect := 22
	runSample(t, pieces, positions, expect)
}

func TestSample3(t *testing.T) {
	pieces := []string{"bishop"}
	positions := [][]int{{4, 3}}
	expect := 12
	runSample(t, pieces, positions, expect)
}

func TestSample4(t *testing.T) {
	pieces := []string{"rook", "rook"}
	positions := [][]int{{1, 1}, {8, 8}}
	expect := 223
	runSample(t, pieces, positions, expect)
}

func TestSample5(t *testing.T) {
	pieces := []string{"queen", "bishop"}
	positions := [][]int{{5, 7}, {3, 4}}
	expect := 281
	runSample(t, pieces, positions, expect)
}
