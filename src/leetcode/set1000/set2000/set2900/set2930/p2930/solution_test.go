package p2930

import "testing"

func runSample(t *testing.T, n int, limit int, expect int) {
	res := int(distributeCandies(n, limit))

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, limit := 5, 2
	expect := 3
	runSample(t, n, limit, expect)
}

func TestSample2(t *testing.T) {
	n, limit := 3, 3
	expect := 10
	runSample(t, n, limit, expect)
}
