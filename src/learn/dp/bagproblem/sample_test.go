package main

import "testing"

var tests = []struct {
	ws     []int
	vs     []int
	weight int
	value  int
}{
	{[]int{5, 10, 20}, []int{50, 60, 140}, 30, 200},
	{[]int{2, 5, 10, 5}, []int{40, 30, 50, 10}, 16, 90},
}

func TestDp(t *testing.T) {
	for _, test := range tests {
		got := dp(test.ws, test.vs, test.weight)
		if got != test.value {
			t.Errorf("maximum value %d from (%v, %v) is not correct: %d\n", got, test.ws, test.vs, test.value)
		}
	}
}
