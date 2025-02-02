package p3442

import "testing"

func runSample(t *testing.T, eventTime int, start []int, end []int, expect int) {
	res := maxFreeTime(eventTime, start, end)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	eventTime := 5
	start := []int{1, 3}
	end := []int{2, 5}
	expect := 2
	runSample(t, eventTime, start, end, expect)
}

func TestSample2(t *testing.T) {
	eventTime := 10
	start := []int{0, 7, 9}
	end := []int{1, 8, 10}
	expect := 7
	runSample(t, eventTime, start, end, expect)
}

func TestSample3(t *testing.T) {
	eventTime := 10
	start := []int{0, 3, 7, 9}
	end := []int{1, 4, 8, 10}
	expect := 6
	runSample(t, eventTime, start, end, expect)
}
