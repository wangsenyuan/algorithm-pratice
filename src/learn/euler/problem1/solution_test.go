package problem1

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := MultipleOf3Or5(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 23)
}

func TestSample2(t *testing.T) {
	runSample(t, 1000, 233168)
}
