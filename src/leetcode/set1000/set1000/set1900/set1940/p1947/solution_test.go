package p1947

import "testing"

func runSample(t *testing.T, students [][]int, mentors [][]int, expect int) {
	res := maxCompatibilitySum(students, mentors)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	students := [][]int{
		{1, 1, 0}, {1, 0, 1}, {0, 0, 1},
	}
	mentors := [][]int{
		{1, 0, 0}, {0, 0, 1}, {1, 1, 0},
	}
	expect := 8
	runSample(t, students, mentors, expect)
}

func TestSample2(t *testing.T) {
	students := [][]int{
		{0, 0}, {0, 0}, {0, 0},
	}
	mentors := [][]int{
		{1, 1}, {1, 1}, {1, 1},
	}
	expect := 0
	runSample(t, students, mentors, expect)
}
