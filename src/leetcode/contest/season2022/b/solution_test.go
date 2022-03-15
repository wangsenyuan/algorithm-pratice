package b

import "testing"

func TestSample1(t *testing.T) {
	sys := Constructor()
	sys.AddActivity(1, 10, 6, 3, 2)
	sys.AddActivity(2, 15, 8, 8, 2)
	res := sys.Consume(101, 13)
	if res != 7 {
		t.Fatalf("Sample expect 7, but got %d", res)
	}
	sys.RemoveActivity(2)
	res = sys.Consume(101, 17)
	if res != 11 {
		t.Fatalf("Sample expect 11, but got %d", res)
	}

	res = sys.Consume(101, 11)
	if res != 11 {
		t.Fatalf("Sample expect 11, but got %d", res)
	}

	res = sys.Consume(102, 16)
	if res != 10 {
		t.Fatalf("Sample expect 10, but got %d", res)
	}

	res = sys.Consume(102, 21)
	if res != 21 {
		t.Fatalf("Sample expect 21, but got %d", res)
	}
}
