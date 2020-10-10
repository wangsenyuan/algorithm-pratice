package sumbyfactors

import "testing"

func runSample(t *testing.T, arr []int, expect string) {
	res := SumOfDivided(arr)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{12, 15}
	expect := "(2 12)(3 27)(5 15)"
	runSample(t, arr, expect)
}

func TestSample2(t *testing.T) {
	arr := []int{15, 30, -45}
	expect := "(2 30)(3 0)(5 0)"
	runSample(t, arr, expect)
}
