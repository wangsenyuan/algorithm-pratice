package p155

import "testing"

func TestSample1(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	if stack.GetMin() != -3 {
		t.Errorf("first getMin should return -3, but got %d", stack.GetMin())
	}
	stack.Pop()
	if stack.Top() != 0 {
		t.Errorf("this time, top should return 0, but got %d", stack.Top())
	}
	if stack.GetMin() != -2 {
		t.Errorf("second getMin should return -2, but got %d", stack.GetMin())
	}
}
