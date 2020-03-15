package p1381

import "testing"

func TestSample1(t *testing.T) {
	stack := Constructor(3)

	stack.Push(1)
	stack.Push(2)

	res := stack.Pop()

	if res != 2 {
		t.Fatalf("expect 2, but got %d", res)
	}

	stack.Push(2)
	stack.Push(3)

	stack.Push(4)

	stack.Increment(5, 100)
	stack.Increment(2, 100)

	res = stack.Pop()

	if res != 103 {
		t.Fatalf("expect 103, but got %d", res)
	}

	res = stack.Pop()

	if res != 202 {
		t.Fatalf("expect 202, but got %d", res)
	}

	res = stack.Pop()

	if res != 201 {
		t.Fatalf("expect 201, but got %d", res)
	}

	res = stack.Pop()

	if res != -1 {
		t.Fatalf("expect -1, but got %d", res)
	}
}
