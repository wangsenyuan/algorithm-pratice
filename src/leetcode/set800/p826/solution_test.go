package p826

import "testing"

func runSample(t *testing.T, difficulty []int, profit []int, worker []int, expect int) {
	res := maxProfitAssignment(difficulty, profit, worker)
	if res != expect {
		t.Errorf("expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	difficulty := []int{2, 4, 6, 8, 10}
	profit := []int{10, 20, 30, 40, 50}
	worker := []int{4, 5, 6, 7}

	expect := 100
	runSample(t, difficulty, profit, worker, expect)
}

func TestSample2(t *testing.T) {
	diff := []int{68, 35, 52, 47, 86}
	prof := []int{67, 17, 1, 81, 3}
	worker := []int{92, 10, 85, 84, 82}
	expect := 324
	runSample(t, diff, prof, worker, expect)
}
