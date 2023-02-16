package main

import "testing"

func runSample(t *testing.T, k int, P []string, S []string, T []int, expect bool) {

	res := solve(k, P, S, T)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	for i := 0; i < len(S); i++ {
		var j int
		for res[j] != T[i] {
			if canMatch(P[res[j]-1], S[i]) {
				t.Fatalf("%s should not match a pattern before %d", S[i], T[i])
			}
			j++
		}
		if res[j] != T[i] || !canMatch(P[res[j]-1], S[i]) {
			t.Fatalf("%s %s, not match", P[res[j]-1], S[i])
		}
	}
}

func canMatch(p, s string) bool {
	for i := 0; i < len(p); i++ {
		if p[i] != s[i] && p[i] != '_' {
			return false
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	k := 4
	P := []string{
		"_b_d",
		"__b_",
		"aaaa",
		"ab__",
		"_bcd",
	}
	S := []string{
		"abcd",
		"abba",
		"dbcd",
	}
	T := []int{4, 2, 5}
	expect := true
	runSample(t, k, P, S, T, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	P := []string{
		"__c",
	}
	S := []string{
		"cba",
	}
	T := []int{1}
	expect := false
	runSample(t, k, P, S, T, expect)
}

func TestSample3(t *testing.T) {
	k := 2
	P := []string{
		"a_", "_b",
	}
	S := []string{
		"ab", "ab",
	}
	T := []int{1, 2}
	expect := false
	runSample(t, k, P, S, T, expect)
}

func TestSample4(t *testing.T) {
	k := 1
	P := []string{
		"_", "b", "c", "a",
	}
	S := []string{
		"c", "a", "b", "a", "b", "b", "b", "a", "b", "c",
	}
	T := []int{1, 4, 2, 4, 3, 1, 3, 3, 1, 1}
	expect := false
	runSample(t, k, P, S, T, expect)
}
