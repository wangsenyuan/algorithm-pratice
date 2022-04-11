package p2234

import "testing"

func runSample(t *testing.T, flowers []int, newFlowers int64, target int, full int, partial int, expect int64) {
	res := maximumBeauty(flowers, newFlowers, target, full, partial)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	flowers := []int{1, 3, 1, 1}
	newFlowers := 7
	target := 6
	full := 12
	partial := 1
	var expect int64 = 14
	runSample(t, flowers, int64(newFlowers), target, full, partial, expect)
}

func TestSample2(t *testing.T) {
	flowers := []int{2, 4, 5, 3}
	newFlowers := 10
	target := 5
	full := 2
	partial := 6
	var expect int64 = 30
	runSample(t, flowers, int64(newFlowers), target, full, partial, expect)
}

func TestSample3(t *testing.T) {
	flowers := []int{8, 20, 13}
	newFlowers := 12
	target := 12
	full := 4
	partial := 3
	var expect int64 = 41
	runSample(t, flowers, int64(newFlowers), target, full, partial, expect)
}

func TestSample4(t *testing.T) {
	flowers := []int{10, 9, 16, 14, 6, 5, 11, 12, 17, 2, 11, 15, 1}
	newFlowers := 80
	target := 14
	full := 15
	partial := 1
	var expect int64 = 195
	runSample(t, flowers, int64(newFlowers), target, full, partial, expect)
}
func TestSample5(t *testing.T) {
	flowers := []int{20, 1, 15, 17, 10, 2, 4, 16, 15, 11}
	newFlowers := 2
	target := 20
	full := 10
	partial := 2
	var expect int64 = 14
	runSample(t, flowers, int64(newFlowers), target, full, partial, expect)
}
