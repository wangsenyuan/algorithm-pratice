package p2940

import "testing"

func runSample(t *testing.T, s1, s2, s3 string, expect int) {
	res := findMinimumOperations(s1, s2, s3)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s1 := "abc"
	s2 := "abb"
	s3 := "ab"
	expect := 2
	runSample(t, s1, s2, s3, expect)
}
