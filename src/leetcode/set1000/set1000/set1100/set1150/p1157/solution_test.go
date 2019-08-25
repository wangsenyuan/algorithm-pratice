package p1157

import "testing"

func TestSample1(t *testing.T) {
	solution := Constructor([]int{1, 1, 2, 2, 1, 1})
	res := solution.Query(0, 5, 4)
	if res != 1 {
		t.Fatalf("expect 1 but got %d at query 0", res)
	}
	res = solution.Query(0, 3, 3)
	if res != -1 {
		t.Fatalf("expect -1, but got %d at query 1", res)
	}

	res = solution.Query(2, 3, 2)
	if res != 2 {
		t.Fatalf("expect 2, but got %d at query 2", res)
	}
}
