package lcp42

import "testing"

func runSample(t *testing.T, toys [][]int, circles [][]int, r int, expect int) {
	res := circleGame(toys, circles, r)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	toys := [][]int{
		{3, 3, 1}, {3, 2, 1},
	}
	circles := [][]int{
		{4, 3},
	}
	r := 2
	expect := 1
	runSample(t, toys, circles, r, expect)
}

func TestSample2(t *testing.T) {
	toys := [][]int{
		{1, 3, 2}, {4, 3, 1}, {7, 1, 2},
	}
	circles := [][]int{
		{1, 0}, {3, 3},
	}
	r := 4
	expect := 2
	runSample(t, toys, circles, r, expect)
}

func TestSample3(t *testing.T) {
	toys := [][]int{
		{3, 4, 5}, {1, 4, 4}, {4, 4, 1}, {1, 5, 5},
	}
	circles := [][]int{
		{4, 1}, {4, 2},
	}
	r := 6
	expect := 1
	runSample(t, toys, circles, r, expect)
}
