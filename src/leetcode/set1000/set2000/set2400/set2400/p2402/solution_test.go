package p2402

import "testing"

func runSample(t *testing.T, n int, meetings [][]int, expect int) {
	res := mostBooked(n, meetings)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	meetings := [][]int{
		{0, 10}, {1, 5}, {2, 7}, {3, 4},
	}
	expect := 0
	runSample(t, n, meetings, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	meetings := [][]int{
		{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8},
	}
	expect := 1
	runSample(t, n, meetings, expect)
}

