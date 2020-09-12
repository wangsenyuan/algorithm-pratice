package b

import "testing"

func runSample(t *testing.T, staple []int, drinks []int, x int, expect int) {
	res := breakfastNumber(staple, drinks, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	staple := []int{10, 20, 5}
	drinks := []int{5, 5, 2}
	expect := 6
	x := 15
	runSample(t, staple, drinks, x, expect)
}

func TestSample2(t *testing.T) {
	staple := []int{2, 1, 1}
	drinks := []int{8, 9, 5, 1}
	expect := 8
	x := 9
	runSample(t, staple, drinks, x, expect)
}
