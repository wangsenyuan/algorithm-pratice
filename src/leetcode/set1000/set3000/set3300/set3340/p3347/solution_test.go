package p3347

import "testing"

func runSample(t *testing.T, a []int, k int, cnt int, expect int) {
	res := maxFrequency(a, k, cnt)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 4, 5}
	k := 1
	cnt := 2
	expect := 2
	runSample(t, a, k, cnt, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, 11, 20, 20}
	k := 5
	cnt := 1
	expect := 2
	runSample(t, a, k, cnt, expect)
}
