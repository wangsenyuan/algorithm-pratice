package main

import "testing"

func runSample(t *testing.T, necklaces []string, k int, expect int) {
	cnt, res := solve(necklaces, k)

	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d, %v", expect, cnt, res)
	}

	if expect < 0 {
		return
	}
	n := len(necklaces)

	arr := make([]string, n)
	copy(arr, necklaces)

	for _, x := range res {
		arr[x-1] = reverse(arr[x-1])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !check(arr[i], arr[j], k) {
				t.Fatalf("Sample result %v, not correct at %d %d", res, i, j)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	k := 2
	necklaces := []string{
		"1010100",
		"0010101",
		"1111010",
		"1000010",
		"0000101",
	}
	expect := 2
	runSample(t, necklaces, k, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	necklaces := []string{
		"011111110",
		"100111000",
		"111100000",
		"000111111",
		"110100111",
		"111110111",
	}
	expect := 1
	runSample(t, necklaces, k, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	necklaces := []string{
		"0001",
		"1000",
		"0000",
	}
	expect := 0
	runSample(t, necklaces, k, expect)
}

func TestSample4(t *testing.T) {
	k := 4
	necklaces := []string{
		"0001",
		"1000",
		"0000",
	}
	expect := -1
	runSample(t, necklaces, k, expect)
}
