package p3002

import "testing"

func runSample(t *testing.T, s string, k int, expect int) {
	res := maxPartitionsAfterOperations(s, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "accca"
	k := 2
	expect := 3
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "aabaab"
	k := 3
	expect := 1
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "baac"
	k := 3
	expect := 2
	runSample(t, s, k, expect)
}

func TestSample4(t *testing.T) {
	s := "aabcacc"
	k := 2
	expect := 3
	runSample(t, s, k, expect)
}

func TestSample5(t *testing.T) {
	s := "aca"
	k := 2
	expect := 2
	runSample(t, s, k, expect)
}

func TestSample6(t *testing.T) {
	s := "acb"
	k := 2
	expect := 2
	runSample(t, s, k, expect)
}

func TestSample7(t *testing.T) {
	s := "baacccb"
	k := 1
	expect := 6
	runSample(t, s, k, expect)
}

func TestSample8(t *testing.T) {
	s := "bccaabcaa"
	k := 2
	expect := 4
	runSample(t, s, k, expect)
}
