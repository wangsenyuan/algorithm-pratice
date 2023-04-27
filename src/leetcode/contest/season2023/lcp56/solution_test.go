package lcp56

import "testing"

func runSample(t *testing.T, mat []string, start []int, end []int, expect int) {
	res := conveyorBelt(mat, start, end)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := []string{
		">>v",
		"v^<",
		"<><",
	}
	start := []int{0, 1}
	end := []int{2, 0}
	expect := 1
	runSample(t, mat, start, end, expect)
}

func TestSample2(t *testing.T) {
	mat := []string{
		">>v", ">>v", "^<<",
	}
	start := []int{0, 0}
	end := []int{1, 1}
	expect := 0
	runSample(t, mat, start, end, expect)
}

func TestSample3(t *testing.T) {
	mat := []string{
		"^v<<",
		">v>^",
		">^>v",
	}
	start := []int{1, 2}
	end := []int{2, 1}
	expect := 0
	runSample(t, mat, start, end, expect)
}

func TestSample4(t *testing.T) {
	mat := []string{
		">^^>",
		"<^v>",
		"^v^<",
	}
	start := []int{0, 0}
	end := []int{1, 3}
	expect := 3
	runSample(t, mat, start, end, expect)
}
