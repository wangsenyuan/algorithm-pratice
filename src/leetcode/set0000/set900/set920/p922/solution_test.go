package p922

import "testing"

func runSample(t *testing.T, A []int) {
	res := sortArrayByParityII(A)
	for i := 0; i < len(res); i++ {
		num := res[i]
		if num&1 == 1 != (i&1 == 1) {
			t.Errorf("Sample %v, got %v, wrong at %d", A, res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{4, 2, 5, 7})
}

func TestSample2(t *testing.T) {
	runSample(t, []int{2, 3})
}
