package p3404

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := numberOfSubsequences(a)

	if res != int64(expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4, 3, 6, 1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 4, 3, 4, 3, 4, 3, 4}
	expect := 3
	runSample(t, a, expect)
}
