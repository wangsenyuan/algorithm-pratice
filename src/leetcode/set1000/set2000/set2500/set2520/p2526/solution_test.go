package p2526

import "testing"

func TestSample1(t *testing.T) {
	sys := Constructor(4, 3)

	res := sys.Consec(4)
	if res {
		t.Fatalf("Sample expect false, but got %t", res)
	}

	res = sys.Consec(4)

	if res {
		t.Fatalf("Sample expect false, but got %t", res)
	}

	res = sys.Consec(4)

	if !res {
		t.Fatalf("Sample expect true, but got %t", res)
	}
}
