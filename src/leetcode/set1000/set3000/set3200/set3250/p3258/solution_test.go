package p3258

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := maxEnergyBoost(a, b)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 1}
	b := []int{3, 1, 1}
	expect := 5
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 1, 1}
	b := []int{1, 1, 3}
	expect := 7
	runSample(t, a, b, expect)
}
