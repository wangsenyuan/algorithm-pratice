package p1894

import "testing"

func runSample(t *testing.T, chalks []int, k int, expect int) {
	res := chalkReplacer(chalks, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	chalks := []int{5, 1, 5}
	k := 22
	expect := 0
	runSample(t, chalks, k, expect)
}

func TestSample2(t *testing.T) {
	chalks := []int{3, 4, 1, 2}
	k := 25
	expect := 1
	runSample(t, chalks, k, expect)
}
