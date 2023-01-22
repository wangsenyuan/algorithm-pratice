package p2542

import "testing"

func runSample(t *testing.T, A, B []int, k int, expect int) {
	res := maxScore(A, B, k)
	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{4, 2, 3, 1, 1}
	B := []int{7, 5, 10, 9, 6}
	k := 1
	expect := 30
	runSample(t, A, B, k, expect)
}
