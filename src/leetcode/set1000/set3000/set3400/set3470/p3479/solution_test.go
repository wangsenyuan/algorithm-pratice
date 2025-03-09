package p3479

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	ans := numOfUnplacedFruits(a, b)
	if ans != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 2, 5}
	b := []int{3, 5, 4}
	expect := 1
	runSample(t, a, b, expect)
}
