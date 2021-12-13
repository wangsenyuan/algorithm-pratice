package p2054

import "testing"

func runSample(t *testing.T, events [][]int, expect int) {
	res := maxTwoEvents(events)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	events := [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}
	expect := 4
	runSample(t, events, expect)
}

func TestSample2(t *testing.T) {
	events := [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}
	expect := 5
	runSample(t, events, expect)
}
