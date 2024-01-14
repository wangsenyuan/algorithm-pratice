package lcr004

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := singleNumber(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 2, 2, 3}
	expect := 3
	runSample(t, a, expect)
}
