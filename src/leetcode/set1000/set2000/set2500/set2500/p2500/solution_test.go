package p2500

import "testing"

func TestSample1(t *testing.T) {
	allo := Constructor(10)

	res := allo.Allocate(1, 1)

	if res != 0 {
		t.Fatalf("Sample expect 0, but got %d", res)
	}

	res = allo.Allocate(1, 2)

	if res != 1 {
		t.Fatalf("Sample expect 1, but got %d", res)
	}

	res = allo.Allocate(1, 3)

	if res != 2 {
		t.Fatalf("Sample expect 2, but got %d", res)
	}

	res = allo.Free(2)

	if res != 1 {
		t.Fatalf("Sample expect 1, but got %d", res)
	}

	res = allo.Allocate(3, 4)

	if res != 3 {
		t.Fatalf("Sample expect 1, but got %d", res)
	}

	res = allo.Allocate(1, 1)

	if res != 1 {
		t.Fatalf("Sample expect 1, but got %d", res)
	}

	res = allo.Allocate(1, 1)

	if res != 6 {
		t.Fatalf("Sample expect 6, but got %d", res)
	}

	res = allo.Free(1)

	if res != 3 {
		t.Fatalf("Sample expect 3, but got %d", res)
	}

	res = allo.Allocate(10, 2)

	if res != -1 {
		t.Fatalf("Sample expect -1, but got %d", res)
	}

	res = allo.Free(7)

	if res != 0 {
		t.Fatalf("Sample expect 0, but got %d", res)
	}
}
