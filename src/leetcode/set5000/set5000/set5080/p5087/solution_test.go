package p5087

import "testing"

func runSample(t *testing.T, tiles string, expect int) {
	res := numTilePossibilities(tiles)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", tiles, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "AAB", 8)
}

func TestSample2(t *testing.T) {
	runSample(t, "AAABBC", 188)
}
