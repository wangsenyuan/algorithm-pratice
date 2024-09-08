package p3280

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := findMaximumScore(a)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 1, 5}
	expect := 7
	runSample(t, a, expect)
}
func TestSample2(t *testing.T) {
	a := []int{4, 3, 1, 3, 2}
	expect := 16
	runSample(t, a, expect)
}
