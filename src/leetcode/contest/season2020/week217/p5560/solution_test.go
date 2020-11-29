package p5560

import "testing"

func TestSample1(t *testing.T) {
	q := Constructor()

	q.PushFront(1)
	q.PushBack(2)
	q.PushMiddle(3)
	q.PushMiddle(4)

	res := q.PopFront()

	if res != 1 {
		t.Fatalf("Sample should get 1, but got %d", res)
	}
	res = q.PopMiddle()
	if res != 3 {
		t.Fatalf("should get 3, but got %d", res)
	}

	res = q.PopMiddle()

	if res != 4 {
		t.Fatalf("should get 4, but got %d", res)
	}

	res = q.PopBack()

	if res != 2 {
		t.Fatalf("Should get 2, but got %d", res)
	}

	res = q.PopFront()

	if res != -1 {
		t.Fatalf("should get -1, but got %d", res)
	}
}
