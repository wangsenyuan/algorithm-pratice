package main

import "testing"

func runSample(t *testing.T, H []int, W []int, expect []int) {
	res := solve(H, W)
	n := len(H)
	for i := 0; i < n; i++ {
		if (expect[i] < 0) != (res[i] < 0) {
			t.Fatalf("Sample expect %d, but got %d at %d", expect[i], res[i], i)
		}
		if res[i] > 0 {
			j := res[i]
			j--
			if H[j] < H[i] && W[j] < W[i] || H[j] < W[i] && W[j] < H[i] {
				continue
			}
			t.Fatalf("Sample result %d !< %d", j, i)
		}
	}
}

func TestSample1(t *testing.T) {
	H := []int{3, 5, 3}
	W := []int{4, 4, 3}
	expect := []int{-1, 3, -1}
	runSample(t, H, W, expect)
}

func TestSample2(t *testing.T) {
	H := []int{1, 2, 3}
	W := []int{3, 2, 1}
	expect := []int{-1, -1, -1}
	runSample(t, H, W, expect)
}

func TestSample3(t *testing.T) {
	H := []int{2, 3, 6, 5}
	W := []int{2, 1, 3, 4}
	expect := []int{-1, -1, 2, 2}
	runSample(t, H, W, expect)
}

func TestSample4(t *testing.T) {
	H := []int{2, 2, 1, 4}
	W := []int{2, 3, 1, 4}
	expect := []int{3, 3, -1, 3}
	runSample(t, H, W, expect)
}

func TestSample5(t *testing.T) {
	H := []int{5, 5}
	W := []int{6, 2}
	expect := []int{2, -1}
	runSample(t, H, W, expect)
}

func TestSample6(t *testing.T) {
	H := []int{4, 9, 5, 10, 9, 10, 9, 1, 6}
	W := []int{7, 2, 10, 8, 2, 7, 8, 10, 3}
	expect := []int{9, -1, 1, 2, -1, 2, 9, -1, -1}
	runSample(t, H, W, expect)
}
