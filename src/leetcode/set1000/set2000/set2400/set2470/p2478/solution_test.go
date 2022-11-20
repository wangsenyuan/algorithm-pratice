package p2478

import "testing"

func runSample(t *testing.T, s string, k int, minLength int, expect int) {
	res := beautifulPartitions(s, k, minLength)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "23542185131"
	k := 3
	minLength := 2
	expect := 3
	runSample(t, s, k, minLength, expect)
}

func TestSample2(t *testing.T) {
	s := "242538614532395749146912679859"
	k := 1
	minLength := 6
	expect := 1
	runSample(t, s, k, minLength, expect)
}
