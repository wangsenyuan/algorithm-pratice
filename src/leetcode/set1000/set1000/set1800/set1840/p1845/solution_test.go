package p1845

import "testing"

func TestSample1(t *testing.T) {
	man := Constructor(5)
	res := man.Reserve()
	if res != 1 {
		t.Fatalf("should get 1, but got %d", res)
	}
	res = man.Reserve()
	if res != 2 {
		t.Fatalf("should get 2, but got %d", res)
	}

	man.Unreserve(2)

	res = man.Reserve()

	if res != 2 {
		t.Fatalf("should get 2, but got %d", res)
	}

	res = man.Reserve()

	if res != 3 {
		t.Fatalf("should get 3, but got %d", res)
	}

	res = man.Reserve()

	if res != 4 {
		t.Fatalf("should get 4, but got %d", res)
	}

	res = man.Reserve()

	if res != 5 {
		t.Fatalf("should get 5, but got %d", res)
	}

	man.Unreserve(5)
}
