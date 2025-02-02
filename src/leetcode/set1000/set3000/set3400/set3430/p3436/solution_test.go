package p3436

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := maxDifference(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	s := "aaaaabbc"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "yzyyys"
	expect := -3
	runSample(t, s, expect)
}
