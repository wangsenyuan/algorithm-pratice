package lcp33

import "testing"

func runSample(t *testing.T, bucket, vat []int, expect int) {
	res := storeWater(bucket, vat)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	bucket := []int{9, 0, 1}
	vat := []int{0, 2, 2}
	expect := 3
	runSample(t, bucket, vat, expect)
}

func TestSample2(t *testing.T) {
	bucket := []int{3, 2, 0}
	vat := []int{0, 0, 0}
	expect := 0
	runSample(t, bucket, vat, expect)
}
