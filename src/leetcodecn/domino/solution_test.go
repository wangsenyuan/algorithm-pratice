package domino

import "testing"

func runSample(t *testing.T, n int, m int, broken [][]int, expect int) {
	res := domino(n, m, broken)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, broken, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 3
	broken := [][]int{
		{1, 0}, {1, 1},
	}
	expect := 2
	runSample(t, n, m, broken, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 3
	broken := [][]int{}
	expect := 4
	runSample(t, n, m, broken, expect)
}

func TestSample4(t *testing.T) {
	n, m := 4, 4
	broken := [][]int{{1, 1}, {2, 2}}
	expect := 6
	runSample(t, n, m, broken, expect)
}

func TestSample5(t *testing.T) {
	n, m := 8, 8
	broken := [][]int{{1, 0}, {2, 5}, {3, 1}, {3, 2}, {3, 4}, {4, 0}, {4, 3}, {4, 6}, {4, 7}, {5, 3}, {5, 5}, {5, 6}, {6, 3}, {7, 2}, {7, 7}}
	expect := 23
	runSample(t, n, m, broken, expect)
}
