package p3438

import (
	"testing"
)

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := minimumIncrements(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4}
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{8, 4}
	b := []int{10, 5}
	expect := 2
	runSample(t, a, b, expect)
}
func TestSample3(t *testing.T) {
	a := []int{7, 9}
	b := []int{7}
	expect := 0
	runSample(t, a, b, expect)
}
