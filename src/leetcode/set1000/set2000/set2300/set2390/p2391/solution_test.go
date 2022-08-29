package p2391

import "testing"

func runSample(t *testing.T, garbages []string, travel []int, expect int) {
	res := garbageCollection(garbages, travel)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	garbages := []string{"G", "P", "GP", "GG"}
	travel := []int{2, 4, 3}
	expect := 21
	runSample(t, garbages, travel, expect)
}
func TestSample2(t *testing.T) {
	garbages := []string{"MMM", "PGM", "GP"}
	travel := []int{3, 10}
	expect := 37
	runSample(t, garbages, travel, expect)
}
