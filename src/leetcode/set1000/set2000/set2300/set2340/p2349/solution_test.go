package p2349

import "testing"

func TestSample1(t *testing.T) {
	containers := Constructor()

	res := containers.Find(10)

	if res != -1 {
		t.Fatalf("Sample expect -1, but got %d", res)
	}

	containers.Change(1000000000, 10)

	res = containers.Find(10)

	if res != 1000000000 {
		t.Fatalf("Sample expect 1000000000, but got %d", res)
	}
}
