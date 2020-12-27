package p1699

import "testing"

func runSample(t *testing.T, a, b []int, expect int) {
	res := countStudents(a, b)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 0, 0}
	b := []int{0, 1, 0, 1}
	expect := 0
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1, 0, 0, 1}
	b := []int{1, 0, 0, 0, 1, 1}
	expect := 3
	runSample(t, a, b, expect)
}
