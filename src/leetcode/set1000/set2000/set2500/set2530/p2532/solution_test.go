package p2532

import "testing"

func runSample(t *testing.T, n int, k int, time [][]int, expect int) {
	res := findCrossingTime(n, k, time)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	k := 3
	time := [][]int{{1, 1, 2, 1}, {1, 1, 3, 1}, {1, 1, 4, 1}}
	expect := 6
	runSample(t, n, k, time, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	k := 2
	time := [][]int{{1, 9, 1, 8}, {10, 10, 10, 10}}
	expect := 50
	runSample(t, n, k, time, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	k := 5
	time := [][]int{{10, 8, 4, 7}, {10, 6, 9, 10}, {3, 7, 4, 6}, {7, 3, 9, 10}, {10, 6, 3, 5}}
	expect := 133
	runSample(t, n, k, time, expect)
}

func TestSample4(t *testing.T) {
	n := 10
	k := 6
	time := [][]int{{2, 10, 5, 8}, {3, 5, 2, 2}, {5, 8, 10, 10}, {7, 8, 8, 5}, {5, 6, 6, 10}, {6, 10, 6, 2}}
	expect := 149
	runSample(t, n, k, time, expect)
}
