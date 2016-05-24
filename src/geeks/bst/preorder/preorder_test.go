package preorder

import "testing"

func TestFastIsPreOrder(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{40, 30, 35, 80, 100}, true},
		{[]int{40, 30, 32, 35, 80, 90, 100, 120}, true},
		{[]int{7, 9, 6, 1, 4, 2, 3, 40}, false},
	}

	for _, test := range tests {
		got := fastIsPreOrder(test.nums)
		if got != test.want {
			t.Errorf(`fastIsPreOrder(%v) == %v`, test.nums, got)
		}
	}
}

func TestSlowIsPreOrder(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{40, 30, 35, 80, 100}, true},
		{[]int{40, 30, 32, 35, 80, 90, 100, 120}, true},
		{[]int{7, 9, 6, 1, 4, 2, 3, 40}, false},
	}

	for _, test := range tests {
		got := slowIsPreOrder(test.nums)
		if got != test.want {
			t.Errorf(`fastIsPreOrder(%v) == %v`, test.nums, got)
		}
	}
}
