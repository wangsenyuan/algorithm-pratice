package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	var cnt int
	for i := 0; i < len(s); i++ {
		if expect[i] == '0' {
			cnt++
		}
		if res[i] == '0' {
			cnt--
		}
	}
	if cnt != 0 {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	for l := 0; l < len(s); l++ {
		for r := l; r < len(s); r++ {
			x := get(s[l : r+1])
			y := get(res[l : r+1])
			if x != y {
				t.Fatalf("Sample expect %s, but got %s", expect, res)
			}
		}
	}
}

func get(s string) int {
	dp := make([]int, 2)
	for _, x := range []byte(s) {
		a, b := dp[0], dp[1]
		if x == '0' {
			dp[0] = a + 1
			dp[1] = b
		} else {
			dp[0] = a + 1
			dp[1] = max(a, b) + 1
		}
	}
	return max(dp[0], dp[1])
}

func TestSample1(t *testing.T) {
	s := "110"
	expect := "010"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "010"
	expect := "010"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "0001111"
	expect := "0000000"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "0111001100111011101000"
	x := "0011001100001011101000"
	runSample(t, s, x)
}
