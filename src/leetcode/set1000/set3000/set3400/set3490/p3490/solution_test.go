package p3490

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := beautifulNumbers(l, r)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 20, 2)
}
