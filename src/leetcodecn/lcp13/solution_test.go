package lcp13

import "testing"

func runSample(t *testing.T, maze []string, expect int) {
	res := minimalSteps(maze)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", maze, expect, res)
	}
}

func TestSample1(t *testing.T) {
	maze := []string{
		"S#O", "M..", "M.T",
	}
	expect := 16
	runSample(t, maze, expect)
}

func TestSample2(t *testing.T) {
	maze := []string{
		"S#O", "M.#", "M.T",
	}
	expect := -1
	runSample(t, maze, expect)
}

func TestSample3(t *testing.T) {
	maze := []string{
		"S#O", "M.T", "M..",
	}
	expect := 17
	runSample(t, maze, expect)
}

func TestSample4(t *testing.T) {
	maze := []string{
		"MMMMM", "MS#MM", "MM#TO",
	}
	expect := 95
	runSample(t, maze, expect)
}
