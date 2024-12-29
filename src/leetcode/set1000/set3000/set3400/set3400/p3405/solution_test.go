package p3405

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect int) {
	res := countGoodArrays(n, m, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 1, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2, 2, 6)
}


func TestSample3(t *testing.T) {
	runSample(t, 5, 2, 0, 2)
}
