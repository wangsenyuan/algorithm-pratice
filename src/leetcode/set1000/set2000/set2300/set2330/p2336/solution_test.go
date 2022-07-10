package p2336

import "testing"

func TestSample1(t *testing.T) {
	target := Constructor()

	target.AddBack(2)

	res := target.PopSmallest()

	if res != 1 {
		t.Fatalf("Sample expect 1, but got %d", res)
	}
	res = target.PopSmallest()

	if res != 2 {
		t.Fatalf("Sample expect 2, but got %d", res)
	}

	res = target.PopSmallest()

	if res != 3 {
		t.Fatalf("Sample expect 3, but got %d", res)
	}
	target.AddBack(1)

	res = target.PopSmallest()

	if res != 1 {
		t.Fatalf("Sample expect 1, but got %d", res)
	}

	res = target.PopSmallest()

	if res != 4 {
		t.Fatalf("Sample expect 2, but got %d", res)
	}

	res = target.PopSmallest()

	if res != 5 {
		t.Fatalf("Sample expect 3, but got %d", res)
	}
}
