package p2276

import "testing"

func TestSample1(t *testing.T) {
	target := Constructor()

	target.Add(2, 3)
	target.Add(7, 10)

	res := target.Count()

	if res != 6 {
		t.Fatalf("Sample expect %d, but got %d", 6, res)
	}

	target.Add(5, 8)

	res = target.Count()

	if res != 8 {
		t.Fatalf("Sample expect %d, but got %d", 8, res)
	}
}
