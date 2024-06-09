package p3176

import "testing"

func runSample(t *testing.T, skills []int, k int, expect int) {
	res := findWinningPlayer(skills, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	skills := []int{4, 2, 6, 3, 9}
	k := 2
	expect := 2
	runSample(t, skills, k, expect)
}
