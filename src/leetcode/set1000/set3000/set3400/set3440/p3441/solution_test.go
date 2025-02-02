package p3441

import "testing"

func runSample(t *testing.T, eventTime int, k int, start []int, end []int, expect int) {
	res := maxFreeTime(eventTime, k, start, end)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	eventTime := 5
	k := 1
	startTime := []int{1, 3}
	endTime := []int{2, 5}
	expect := 2
	runSample(t, eventTime, k, startTime, endTime, expect)
}

func TestSample2(t *testing.T) {
	eventTime := 10
	k := 1
	startTime := []int{0, 2, 9}
	endTime := []int{1, 4, 10}
	expect := 6
	runSample(t, eventTime, k, startTime, endTime, expect)
}

func TestSample3(t *testing.T) {
	eventTime := 5
	k := 2
	startTime := []int{0, 1, 2, 3, 4}
	endTime := []int{1, 2, 3, 4, 5}
	expect := 0
	runSample(t, eventTime, k, startTime, endTime, expect)
}
