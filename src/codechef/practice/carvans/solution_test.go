package main

import "testing"

func TestRun(t *testing.T) {
	tests := []struct {
		cars []int
		want int
	}{
		{[]int{10}, 1},
		{[]int{8, 3, 6}, 2},
		{[]int{4, 5, 1, 2, 3}, 2},
	}

	for _, test := range tests {
		got := run(test.cars)
		if got != test.want {
			t.Errorf("%v should have %d at max speed, but get %d\n", test.cars, test.want, got)
		}
	}
}
