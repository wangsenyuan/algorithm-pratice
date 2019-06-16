package p895

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, ops []int, expect []int) {
	stack := Constructor()
	res := make([]int, 0, len(expect))
	for _, op := range ops {
		if op > 0 {
			//push
			stack.Push(op)
		} else {
			ans := stack.Pop()
			res = append(res, ans)
		}
	}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", ops, expect, res)
	}

}

func TestSample1(t *testing.T) {
	ops := []int{5, 7, 5, 7, 4, 5, 0, 0, 0, 0}
	expect := []int{5, 7, 5, 4}
	runSample(t, ops, expect)
}
