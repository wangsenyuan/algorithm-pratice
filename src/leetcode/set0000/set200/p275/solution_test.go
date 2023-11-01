package p275

import "testing"

func runSample(t *testing.T, c []int, expect int) {
	res := hIndex(c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	c := []int{0}
	expect := 0
	runSample(t, c, expect)
}

func TestSample2(t *testing.T) {
	c := []int{1}
	expect := 1
	runSample(t, c, expect)
}
