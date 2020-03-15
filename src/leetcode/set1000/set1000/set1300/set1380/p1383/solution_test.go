package p1383

import "testing"

func runSample(t *testing.T, n int, speed []int, efficiency []int, k int, expect int) {
	res := maxPerformance(n, speed, efficiency, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	speed := []int{2, 10, 3, 1, 5, 8}
	efficiency := []int{5, 4, 3, 9, 7, 2}
	k := 2
	expect := 60
	runSample(t, n, speed, efficiency, k, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	speed := []int{2, 10, 3, 1, 5, 8}
	efficiency := []int{5, 4, 3, 9, 7, 2}
	k := 3
	expect := 68
	runSample(t, n, speed, efficiency, k, expect)
}

func TestSample3(t *testing.T) {
	n := 6
	speed := []int{2, 10, 3, 1, 5, 8}
	efficiency := []int{5, 4, 3, 9, 7, 2}
	k := 4
	expect := 72
	runSample(t, n, speed, efficiency, k, expect)
}

func TestSample4(t *testing.T) {
	n := 3
	speed := []int{2, 8, 2}
	efficiency := []int{2, 7, 1}
	k := 2
	expect := 56
	runSample(t, n, speed, efficiency, k, expect)
}
