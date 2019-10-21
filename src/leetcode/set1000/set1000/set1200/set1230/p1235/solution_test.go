package p1235

import "testing"

func runSample(t *testing.T, startTime []int, endTime []int, profit []int, expect int) {
	res := jobScheduling(startTime, endTime, profit)

	if res != expect {
		t.Errorf("Sample %v %v %v, expect %d, but got %d", startTime, endTime, profit, expect, res)
	}
}

func TestSample1(t *testing.T) {
	startTime := []int{1, 2, 3, 3}
	endTime := []int{3, 4, 5, 6}
	profit := []int{50, 10, 40, 70}
	expect := 120
	runSample(t, startTime, endTime, profit, expect)
}

func TestSample2(t *testing.T) {
	startTime := []int{1, 2, 3, 4, 6}
	endTime := []int{3, 5, 10, 6, 9}
	profit := []int{20, 20, 100, 70, 60}
	expect := 150
	runSample(t, startTime, endTime, profit, expect)
}

func TestSample3(t *testing.T) {
	startTime := []int{1, 1, 1}
	endTime := []int{2, 3, 4}
	profit := []int{5, 6, 4}
	expect := 6
	runSample(t, startTime, endTime, profit, expect)
}
