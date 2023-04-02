package p2611

import "testing"

func runSample(t *testing.T, r1 []int, r2 []int, k int, expect int) {
	res := miceAndCheese(r1, r2, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	r1 := []int{1, 1, 3, 4}
	r2 := []int{4, 4, 1, 1}
	k := 2
	expect := 15
	runSample(t, r1, r2, k, expect)
}

func TestSample2(t *testing.T) {
	r1 := []int{1, 1}
	r2 := []int{1, 1}
	k := 1
	expect := 2
	runSample(t, r1, r2, k, expect)
}
