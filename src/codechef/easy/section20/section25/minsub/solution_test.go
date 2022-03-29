package main

import "testing"

func runSample(t *testing.T, A string, expect int) {
	S, T := solve(A)

	L := longest_common_substring(S, T)
	if L != expect {
		t.Errorf("Sample expect %d, but got %s %s", expect, S, T)
	}
}

func longest_common_substring(a, b string) int {
	var best int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			var k int
			for i+k < len(a) && j+k < len(b) && a[i+k] == b[j+k] {
				k++
			}
			if best < k {
				best = k
			}
		}
	}
	return best
}

func TestSample1(t *testing.T) {
	A := "101"
	expect := 2
	runSample(t, A, expect)
}


func TestSample2(t *testing.T) {
	A := "1111"
	expect := 4
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := "0111"
	expect := 2
	runSample(t, A, expect)
}
