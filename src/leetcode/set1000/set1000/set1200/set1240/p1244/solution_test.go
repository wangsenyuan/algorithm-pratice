package p1244

import "testing"

func TestSample1(t *testing.T) {
	target := Constructor()

	target.AddScore(1, 73)
	target.AddScore(2, 56)
	target.AddScore(3, 39)
	target.AddScore(4, 51)
	target.AddScore(5, 4)
	res := target.Top(1)
	if res != 73 {
		t.Fatalf("Sample Fail at top 1, should get 73, but got %d", res)
	}

	target.Reset(1)
	target.Reset(2)

	target.AddScore(2, 51)
	res = target.Top(3)

	if res != 141 {
		t.Fatalf("Sample Fail at top 3, should get 141, but got %d", res)
	}

	res = target.Top(1)

	if res != 51 {
		t.Fatalf("Sample Fail at top 1, should get 51, but got %d", res)
	}
}
