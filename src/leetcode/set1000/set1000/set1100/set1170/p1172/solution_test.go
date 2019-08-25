package p1172

import "testing"

func TestSample1(t *testing.T) {
	plates := Constructor(2)

	plates.Push(1)
	plates.Push(2)
	plates.Push(3)
	plates.Push(4)
	plates.Push(5)

	res := plates.PopAtStack(0)

	if res != 2 {
		t.Fatalf("Sample fail at step 1, should get 2, but got %d", res)
	}
	plates.Push(20)
	plates.Push(21)
	res = plates.PopAtStack(0)
	if res != 20 {
		t.Fatalf("Sample fail at step 2, should get 20, but got %d", res)
	}
	res = plates.PopAtStack(2)
	if res != 21 {
		t.Fatalf("Sample fail at step 3, should get 21, but got %d", res)
	}

	res = plates.Pop()
	if res != 5 {
		t.Fatalf("Sample fail at step 4, should get 5, but got %d", res)
	}

	res = plates.Pop()
	if res != 4 {
		t.Fatalf("Sample fail at step 5, should get 4, but got %d", res)
	}

	res = plates.Pop()
	if res != 3 {
		t.Fatalf("Sample fail at step 6, should get 3, but got %d", res)
	}

	res = plates.Pop()
	if res != 1 {
		t.Fatalf("Sample fail at step 7, should get 1, but got %d", res)
	}

	res = plates.Pop()
	if res != -1 {
		t.Fatalf("Sample fail at step 8, should get -1, but got %d", res)
	}
}
