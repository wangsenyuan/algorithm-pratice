package p1375

import "testing"

func runSample(t *testing.T, light []int, expect int) {
	res := numTimesAllBlue(light)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", light, expect, res)
	}
}

func TestSample1(t *testing.T) {
	light := []int{2, 1, 3, 5, 4}
	expect := 3
	runSample(t, light, expect)
}

func TestSample2(t *testing.T) {
	light := []int{3, 2, 4, 1, 5}
	expect := 2
	runSample(t, light, expect)
}

func TestSample3(t *testing.T) {
	light := []int{1, 2, 3, 4, 5, 6}
	expect := 6
	runSample(t, light, expect)
}
