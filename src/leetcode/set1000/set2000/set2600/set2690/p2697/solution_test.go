package p2697

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := punishmentNumber(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	expect := 182
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 37
	expect := 1478
	runSample(t, n, expect)
}
