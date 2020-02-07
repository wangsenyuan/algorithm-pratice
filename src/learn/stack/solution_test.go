package stack

import "testing"

func TestIsPopOrder(t *testing.T) {
	push := []int{1, 2, 3, 4, 5}
	pop := []int{4, 5, 3, 2, 1}
	res := isPopOrder(push, pop)

	if !res {
		t.Errorf("Sample %v %v, expect true, but got false", push, pop)
	}
}
