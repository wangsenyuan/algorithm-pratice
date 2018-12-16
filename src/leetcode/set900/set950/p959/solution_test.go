package p959

import "testing"

func runSample(t *testing.T, grid []string, expect int) {
	res := regionsBySlashes(grid)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		" /",
		"/ ",
	}
	runSample(t, grid, 2)
}

func TestSample2(t *testing.T) {
	grid := []string{
		" /",
		"  ",
	}
	runSample(t, grid, 1)
}

func TestSample3(t *testing.T) {
	grid := []string{
		"\\/",
		"/\\",
	}
	runSample(t, grid, 4)
}

func TestSample4(t *testing.T) {
	grid := []string{
		"/\\",
		"\\/",
	}
	runSample(t, grid, 5)
}

func TestSample5(t *testing.T) {
	grid := []string{
		"//",
		"/ ",
	}
	runSample(t, grid, 3)
}
