package p2491

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(dividePlayers(A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 2, 5, 1, 3, 4}
	expect := 22
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 4}
	expect := 12
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 2, 3}
	expect := -1
	runSample(t, A, expect)
}
