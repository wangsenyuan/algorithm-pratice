package p146

import "testing"

func TestSample1(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	res := cache.Get(1)
	if res != 1 {
		t.Fatalf("Sample expect 1, but got %d", res)
	}
	cache.Put(3, 3)
	res = cache.Get(2)
	if res != -1 {
		t.Fatalf("Sample expect -1, but got %d", res)
	}
	cache.Put(4, 4)
	res = cache.Get(1)
	if res != -1 {
		t.Fatalf("Sample expect -1, but got %d", res)
	}
}
