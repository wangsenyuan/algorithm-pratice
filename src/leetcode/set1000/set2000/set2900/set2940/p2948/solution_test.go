package p2948

import "testing"

func runSample(t *testing.T, mat [][]int, k int, expect bool) {
	res := areSimilar(mat, k)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	mat := [][]int{
		{1, 2},
	}
	k := 1
	expect := false
	runSample(t, mat, k, expect)
}
