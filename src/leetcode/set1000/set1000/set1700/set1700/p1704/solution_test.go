package p1704

import "testing"

func runSample(t *testing.T, apples []int, days []int, expect int) {
	res := eatenApples(apples, days)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	apples := []int{1, 2, 3, 5, 2}
	days := []int{3, 2, 1, 4, 2}
	expect := 7
	runSample(t, apples, days, expect)
}

func TestSample2(t *testing.T) {
	apples := []int{3, 0, 0, 0, 0, 2}
	days := []int{3, 0, 0, 0, 0, 2}
	expect := 5
	runSample(t, apples, days, expect)
}

func TestSample3(t *testing.T) {
	apples := []int{1}
	days := []int{2}
	expect := 1
	runSample(t, apples, days, expect)
}
