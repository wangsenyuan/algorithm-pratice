package main

import "testing"

func runSample(t *testing.T, s string, expect []int) {
	res := solve(s)

	a := process(s, expect)
	b := process(s, res)

	if a != b {
		t.Fatalf("Sample result %v => %s, not get correct result %s", res, b, a)
	}
}

func process(s string, expect []int) string {
	buf := []byte(s)

	swap_prefix := func(l int) {
		for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
			buf[i], buf[j] = buf[j], buf[i]
		}
	}

	for i := 0; i < len(s); i++ {
		if expect[i] == 1 {
			swap_prefix(i + 1)
		}
	}

	return string(buf)
}

func TestSample1(t *testing.T) {
	s := "bbab"
	expect := []int{0, 1, 1, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aaaa"
	expect := []int{1, 0, 0, 0, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "abbaabba"
	// abbaabba
	// abbaba
	// bbaaba
	// aabbba
	// abbbaa
	expect := []int{0, 0, 1, 0, 1, 0, 1, 1}
	runSample(t, s, expect)
}
