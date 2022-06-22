package main

import "testing"

func runSample(t *testing.T, K int, V []int, L string, S int, expect string) {
	res := solve(K, V, L, S)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	K := 5
	V := []int{12, 1, 12, 4, 4}
	L := "acccdadceb"
	S := 2
	expect := ""
	runSample(t, K, V, L, S, expect)
}

func TestSample2(t *testing.T) {
	K := 3
	V := []int{5, 4, 6}
	L := "abcbacbabcc"
	S := 15
	expect := "aaa"
	runSample(t, K, V, L, S, expect)
}

func TestSample3(t *testing.T) {
	K := 2
	V := []int{3, 4}
	L := "baba"
	S := 7
	expect := "ab"
	runSample(t, K, V, L, S, expect)
}
