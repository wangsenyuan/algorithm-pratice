package p5557

import "testing"

func runSample(t *testing.T, dest []int, k int, expect string) {
	res := kthSmallestPath(dest, k)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	dest := []int{2, 3}
	k := 1
	expect := "HHHVV"
	runSample(t, dest, k, expect)
}

func TestSample2(t *testing.T) {
	dest := []int{2, 3}
	k := 2
	expect := "HHVHV"
	runSample(t, dest, k, expect)
}

func TestSample3(t *testing.T) {
	dest := []int{2, 3}
	k := 3
	expect := "HHVVH"
	runSample(t, dest, k, expect)
}
