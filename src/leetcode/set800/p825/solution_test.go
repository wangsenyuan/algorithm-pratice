package p825

import "testing"

func runSample(t *testing.T, ages []int, expect int) {
	res := numFriendRequests(ages)

	if res != expect {
		t.Errorf("sample %v, expect %d, but got %d", ages, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{16, 16}, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{16, 17, 18}, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{20, 30, 100, 110, 120}, 3)
}

func TestSample4(t *testing.T) {
	ages := []int{73, 106, 39, 6, 26, 15, 30, 100, 71, 35, 46, 112, 6, 60, 110}
	runSample(t, ages, 29)
}

func TestSample5(t *testing.T) {
	runSample(t, []int{100, 100}, 2)
}
