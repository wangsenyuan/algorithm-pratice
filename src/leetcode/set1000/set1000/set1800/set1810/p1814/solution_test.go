package p1814

import "testing"

func runSample(t *testing.T, batchSize int, groups []int, expect int) {
	res := maxHappyGroups(batchSize, groups)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	batchSize := 4
	groups := []int{1, 3, 2, 5, 2, 2, 1, 6}
	expect := 4
	runSample(t, batchSize, groups, expect)
}
