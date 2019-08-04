package p1146

import "testing"

func TestSample1(t *testing.T) {
	sample := Constructor(3)
	sample.Set(0, 5)
	sample.Snap()
	sample.Set(0, 6)
	res := sample.Get(0, 0)
	if res != 5 {
		t.Errorf("Sample should get 5, but got %d", res)
	}
}

func TestSample2(t *testing.T) {
	sample := Constructor(1)
	sample.Set(0, 15)
	sample.Snap()
	sample.Snap()
	sample.Snap()
	res := sample.Get(0, 2)
	if res != 15 {
		t.Errorf("Sample should get 5, but got %d", res)
	}

	sample.Snap()
	sample.Snap()
	res = sample.Get(0, 0)

	if res != 15 {
		t.Errorf("Sample should get 5, but got %d", res)
	}
}

func TestSample3(t *testing.T) {
	sample := Constructor(4)
	// sample.Set(0, 15)
	sample.Snap()
	sample.Snap()
	// sample.Snap()
	res := sample.Get(3, 1)
	if res != 0 {
		t.Errorf("Sample should get 5, but got %d", res)
	}
	sample.Set(2, 4)

	sample.Snap()

	sample.Set(1, 4)
}
