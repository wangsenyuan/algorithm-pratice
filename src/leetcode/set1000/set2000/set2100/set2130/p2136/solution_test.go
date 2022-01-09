package p2136

import "testing"

func runSample(t *testing.T, plant []int, grow []int, expect int) {
	res := earliestFullBloom(plant, grow)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	plant := []int{1, 4, 3}
	grow := []int{2, 3, 1}
	expect := 9
	runSample(t, plant, grow, expect)
}
