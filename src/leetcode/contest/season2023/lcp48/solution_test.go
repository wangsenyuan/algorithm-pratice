package lcp48

import "testing"

func runSample(t *testing.T, pieces [][]int, expect string) {
	res := gobang(pieces)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	pieces := [][]int{
		{0, 0, 1}, {1, 1, 1}, {2, 2, 0},
	}
	expect := "None"
	runSample(t, pieces, expect)
}

func TestSample2(t *testing.T) {
	pieces := [][]int{
		{1, 2, 1}, {1, 4, 1}, {1, 5, 1}, {2, 1, 0}, {2, 3, 0}, {2, 4, 0}, {3, 2, 1}, {3, 4, 0}, {4, 2, 1}, {5, 2, 1},
	}
	expect := "Black"
	runSample(t, pieces, expect)
}

func TestSample3(t *testing.T) {
	pieces := [][]int{
		{0, 3, 1}, {1, 2, 1}, {1, 3, 1}, {2, 0, 0}, {2, 1, 0}, {2, 4, 0}, {3, 3, 1}, {3, 4, 0}, {4, 3, 1}, {4, 5, 0}, {5, 6, 0},
	}
	expect := "Black"
	runSample(t, pieces, expect)
}

func TestSample4(t *testing.T) {
	pieces := [][]int{
		{0, 0, 1}, {0, 1, 0}, {0, 3, 0}, {0, 4, 0}, {0, 7, 0}, {0, 8, 1},
	}
	expect := "Black"
	runSample(t, pieces, expect)
}

func TestSample5(t *testing.T) {
	pieces := [][]int{
		{0, 0, 1}, {0, 1, 0}, {0, 2, 0}, {0, 3, 0}, {0, 7, 0}, {0, 8, 0}, {0, 9, 0}, {0, 10, 1},
	}
	expect := "Black"
	runSample(t, pieces, expect)
}
