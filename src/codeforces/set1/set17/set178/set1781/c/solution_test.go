package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	_, res := solve(s)

	if !check(res) {
		t.Fatalf("Sample result %s, not correct", res)
	}

	a := countDiff(s, expect)
	b := countDiff(s, res)

	if a != b {
		t.Fatalf("count diff not correct %d vs %d", a, b)
	}
}

func check(s string) bool {
	freq := make([]int, 26)
	var cnt int
	n := len(s)
	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		freq[x]++
		if freq[x] == 1 {
			cnt++
		}
	}

	if n%cnt != 0 {
		return false
	}
	m := n / cnt
	for i := 0; i < 26; i++ {
		if freq[i] == 0 {
			continue
		}
		if freq[i] != m {
			return false
		}
	}

	return true
}

func countDiff(a, b string) int {
	var res int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			res++
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	s := "hello"
	expect := "helno"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "codeforces"
	expect := "codefofced"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "eevee"
	expect := "eeeee"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "appall"
	expect := "appall"
	runSample(t, s, expect)
}
