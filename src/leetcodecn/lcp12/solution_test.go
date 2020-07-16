package lcp12

import "testing"

func runSample(t *testing.T, time []int, m int, expect int) {
	res := minTime(time, m)

	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", time, m, expect, res)
	}
}

func TestSample1(t *testing.T) {
	time := []int{1, 2, 3, 3}
	m := 2
	expect := 3
	runSample(t, time, m, expect)
}

func TestSample2(t *testing.T) {
	time := []int{999, 999, 999}
	m := 4
	expect := 0
	runSample(t, time, m, expect)
}
