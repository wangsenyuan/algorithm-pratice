package p853

import "testing"

func runSample(t *testing.T, target int, position []int, speed []int, expect int) {
	res := carFleet(target, position, speed)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", target, position, speed, expect, res)
	}
}

func TestSample1(t *testing.T) {
	target := 12
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}
	expect := 3
	runSample(t, target, position, speed, expect)
}

func TestSample2(t *testing.T) {
	target := 13
	position := []int{10, 2, 5, 7, 4, 6, 11}
	speed := []int{7, 5, 10, 5, 9, 4, 1}
	expect := 2
	runSample(t, target, position, speed, expect)
}
