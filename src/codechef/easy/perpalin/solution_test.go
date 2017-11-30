package main

import "testing"

func isPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func TestSample(t *testing.T) {
	samples := [][]int{
		{3, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{6, 3},
	}
	expect := []bool{false, false, true, true, true}

	for i := 0; i < len(samples); i++ {
		n, p := samples[i][0], samples[i][1]
		res, can := solve(n, p)
		if can != expect[i] {
			t.Errorf("sample %d %d should %v, but get %v", n, p, expect[i], can)
			continue
		}
		if !can {
			continue
		}
		if !isPalindrome(res) {
			t.Errorf("sample %d %d, result %s should be palindrome, but is not", n, p, res)
		}
		for j := 0; j < p; j++ {
			for k := j; k < n; k += p {
				if res[j] != res[k] {
					t.Errorf("sample %d %d give wrong answer %s, letters at %d & %d differs", n, p, res, j, k)
				}
			}
		}
	}
}
