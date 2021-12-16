package main

import "testing"

func runSample(t *testing.T, n, m int, expect int) {
	res := solve(n, m)

	if len(res) != expect {
		t.Errorf("Sample result %s, not correct, length should be %d", res, expect)
	} else {
		a, b := countHillAndVally(res)
		if a != n || b != m {
			t.Errorf("Sample result %s, not correct, got %d %d", res, a, b)
		}
	}
}

func countHillAndVally(s string) (a int, b int) {
	for i := 1; i < len(s)-1; i++ {
		if s[i] > s[i-1] && s[i] > s[i+1] {
			a++
		} else if s[i] < s[i-1] && s[i] < s[i+1] {
			b++
		}
	}
	return
}

func TestSample1(t *testing.T) {
	n, m := 3, 2
	expect := len("0101010")
	runSample(t, n, m, expect)
}


func TestSample2(t *testing.T) {
	n, m := 2, 3
	expect := len("1010101")
	runSample(t, n, m, expect)
}

func TestSample3(t *testing.T) {
	n, m := 3, 3
	expect := len("01010101")
	runSample(t, n, m, expect)
}
