package p1448

import "testing"

func runSample(t *testing.T, cost []int, target int, expect string) {
	res := largestNumber(cost, target)

	if res != expect {
		t.Errorf("Sample %v %d, expect %s, but got %s", cost, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	cost := []int{4, 3, 2, 5, 6, 7, 2, 5, 5}
	target := 9
	expect := "7772"
	runSample(t, cost, target, expect)
}

func TestSample2(t *testing.T) {
	cost := []int{7, 6, 5, 5, 5, 6, 8, 7, 8}
	target := 12
	expect := "85"
	runSample(t, cost, target, expect)
}

func TestSample3(t *testing.T) {
	cost := []int{2, 4, 6, 2, 4, 6, 4, 4, 4}
	target := 5
	expect := "0"
	runSample(t, cost, target, expect)
}

func TestSample4(t *testing.T) {
	cost := []int{6, 10, 15, 40, 40, 40, 40, 40, 40}
	target := 47
	expect := "32211"
	runSample(t, cost, target, expect)
}
