package p1583

import "testing"

func runSample(t *testing.T, n int, preferences [][]int, pairs [][]int, expect int) {
	res := unhappyFriends(n, preferences, pairs)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	pref := [][]int{
		{1, 2, 3}, {3, 2, 0}, {3, 1, 0}, {1, 2, 0},
	}
	pairs := [][]int{
		{0, 1}, {2, 3},
	}

	expect := 2
	runSample(t, n, pref, pairs, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	pref := [][]int{
		{1}, {0},
	}
	pairs := [][]int{
		{0, 1},
	}

	expect := 0
	runSample(t, n, pref, pairs, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	pref := [][]int{
		{1, 3, 2}, {2, 3, 0}, {1, 3, 0}, {0, 2, 1},
	}
	pairs := [][]int{
		{1, 3}, {0, 2},
	}

	expect := 4
	runSample(t, n, pref, pairs, expect)
}
