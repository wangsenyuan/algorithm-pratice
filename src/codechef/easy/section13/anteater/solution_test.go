package main

import "testing"

func runSample(t *testing.T, R, C int, grid []string, expect int) {
	res := solve(R, C, grid)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", R, C, grid, expect, res)
	}
}

func TestSample1(t *testing.T) {
	R, C := 3, 3
	grid := []string{
		"R--",
		"---",
		"--U",
	}
	expect := 1
	runSample(t, R, C, grid, expect)
}

func TestSample2(t *testing.T) {
	R, C := 1, 4
	grid := []string{
		"R--R",
	}
	expect := 0
	runSample(t, R, C, grid, expect)
}

func TestSample3(t *testing.T) {
	R, C := 1, 4
	grid := []string{
		"R--L",
	}
	expect := 0
	runSample(t, R, C, grid, expect)
}

func TestSample4(t *testing.T) {
	R, C := 1, 4
	grid := []string{
		"-R-L",
	}
	expect := 1
	runSample(t, R, C, grid, expect)
}

func TestSample5(t *testing.T) {
	R, C := 1, 4
	grid := []string{
		"-R#L",
	}
	expect := 0
	runSample(t, R, C, grid, expect)
}

func TestSample6(t *testing.T) {
	R, C := 3, 3
	grid := []string{
		"R-D",
		"-#-",
		"R-U",
	}
	expect := 3
	runSample(t, R, C, grid, expect)
}

func TestSample7(t *testing.T) {
	R, C := 3, 3
	grid := []string{
		"R-D",
		"---",
		"R#U",
	}
	expect := 2
	runSample(t, R, C, grid, expect)
}

func TestSample8(t *testing.T) {
	R, C := 3, 3
	grid := []string{
		"-D-",
		"R-L",
		"-U-",
	}
	expect := 6
	runSample(t, R, C, grid, expect)
}

func TestSample9(t *testing.T) {
	R, C := 1, 7
	grid := []string{
		"RLLLLLL",
	}
	expect := 3
	runSample(t, R, C, grid, expect)
}
