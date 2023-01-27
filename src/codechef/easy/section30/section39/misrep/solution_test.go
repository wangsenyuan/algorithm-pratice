package main

import "testing"

func runSample(t *testing.T, A []int, expect bool) {
	arr := make([]int, len(A))
	copy(arr, A)
	ok, res := solve(A)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !ok {
		return
	}

	for _, op := range res {
		i, j := op[0], op[1]
		i--
		j--
		x := min(arr[i], arr[j])
		arr[i] -= x
		arr[j] -= x
	}
	for _, a := range arr {
		if a > 0 {
			t.Fatalf("Sample result %v, got incorrect array %v", res, arr)
		}
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1}
	expect := true
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 3, 1}
	expect := false
	runSample(t, A, expect)
}
func TestSample3(t *testing.T) {
	A := []int{1, 3, 1, 3}
	expect := true
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 1}
	expect := true
	runSample(t, A, expect)
}

func TestSample5(t *testing.T) {
	A := []int{1, 2, 3}
	expect := true
	runSample(t, A, expect)
}

func TestSample6(t *testing.T) {
	A := []int{1, 2, 3, 4, 8, 2}
	expect := true
	runSample(t, A, expect)
}
