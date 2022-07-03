package p2326

import "testing"

func runSample(t *testing.T, n int, delay int, forget int, expect int) {
	res := peopleAwareOfSecret(n, delay, forget)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	delay := 2
	forget := 4
	expect := 5

	runSample(t, n, delay, forget, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	delay := 1
	forget := 3
	expect := 6

	runSample(t, n, delay, forget, expect)
}
