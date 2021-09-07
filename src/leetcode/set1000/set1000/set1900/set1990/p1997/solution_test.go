package p1997

import "testing"

func runSample(t *testing.T, nextVisit []int, expect int) {
	res := firstDayBeenInAllRooms(nextVisit)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	next := []int{0, 0}
	expect := 2
	runSample(t, next, expect)
}

func TestSample2(t *testing.T) {
	next := []int{0, 0, 2}
	expect := 6
	runSample(t, next, expect)
}

func TestSample3(t *testing.T) {
	next := []int{0, 1, 2, 0}
	expect := 6
	runSample(t, next, expect)
}

func TestSample4(t *testing.T) {
	next := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 8}
	expect := 2048
	runSample(t, next, expect)
}
