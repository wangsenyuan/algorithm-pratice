package p2528

import "testing"

func runSample(t *testing.T, stations []int, r int, k int, expect int64) {
	res := maxPower(stations, r, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	stations := []int{1, 2, 4, 5, 0}
	r := 1
	k := 2
	var expect int64 = 5
	runSample(t, stations, r, k, expect)
}

func TestSample2(t *testing.T) {
	stations := []int{4, 4, 4, 4}
	r := 0
	k := 3
	var expect int64 = 4
	runSample(t, stations, r, k, expect)
}
