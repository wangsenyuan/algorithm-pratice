package p5541

import "testing"

func runSample(tc *testing.T, s, t string, expect int) {
	res := countSubstrings(s, t)

	if res != expect {
		tc.Errorf("Sample %s %s, expect %d, but got %d", s, t, expect, res)
	}
}

func TestSample1(tc *testing.T) {
	s := "aba"
	t := "baba"
	expect := 6
	runSample(tc, s, t, expect)
}

func TestSample2(tc *testing.T) {
	s := "ab"
	t := "bb"
	expect := 3
	runSample(tc, s, t, expect)
}

func TestSample3(tc *testing.T) {
	s := "a"
	t := "a"
	expect := 0
	runSample(tc, s, t, expect)
}

func TestSample4(tc *testing.T) {
	s := "abe"
	t := "bbc"
	expect := 10
	runSample(tc, s, t, expect)
}

func TestCountDigits(tc *testing.T) {
	for i := 0; i < 10000; i++ {
		expect := countDirect(i)

		res := countDigits(i)

		if expect != res {
			tc.Fatalf("%d expect %d, but got %d", i, expect, res)
		}
	}
}

func countDirect(num int) int {
	var cnt int
	for num > 0 {
		cnt += num & 1
		num >>= 1
	}
	return cnt
}
