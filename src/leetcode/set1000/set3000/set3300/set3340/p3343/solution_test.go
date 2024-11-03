package p3343

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := countBalancedPermutations(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "123"
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "112"
	expect := 1
	runSample(t, s, expect)
}
