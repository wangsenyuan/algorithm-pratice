package main

import "testing"

func runSample(t *testing.T, n, k int, ok bool) {
	res := solve(n, k)

	if !ok != (res == "-1") {
		t.Fatalf("Sample expect no answer but got %s", res)
	}

	if !ok {
		return
	}

	if len(res) != n {
		t.Fatalf("Sample result %s, length not correct, %d", res, n)
	}

	var sum int
	for i := 0; i < len(res); i++ {
		j := int(res[i] - 'a')
		sum += 1 << j
	}
	if sum != k {
		t.Fatalf("Sample result %s, not getting the correct score %d", res, k)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 5, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 5, true)
}
