package p992

import "testing"

func runSample(t *testing.T, A []int, K int, expect int) {
	res := subarraysWithKDistinct(A, K)

	if res != expect {
		t.Errorf("Sample %v, %d, expect %d but got %d", A, K, expect, res)
	}
}
func TestSample1(t *testing.T) {
	A := []int{1, 2, 1, 2, 3}
	K := 2
	runSample(t, A, K, 7)
}
