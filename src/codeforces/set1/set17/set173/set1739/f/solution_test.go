package main

import "testing"

func runSample(t *testing.T, C []int, words []string, expect string) {
	res := solve(C, words)

	a := calcTotalCost(C, words, expect)
	b := calcTotalCost(C, words, res)

	if a != b {
		t.Fatalf("Sample expect %s => %d, but got %s => %d", expect, a, res, b)
	}
}

func calcTotalCost(C []int, words []string, keyboard string) int {
	var res int
	for i, word := range words {
		if checkGood(keyboard, word) {
			res += C[i]
		}
	}
	return res
}

func checkGood(keyboard string, word string) bool {
	pos := make([]int, K)
	for i := 0; i < len(keyboard); i++ {
		pos[int(keyboard[i]-'a')] = i
	}
	for i := 0; i+1 < len(word); i++ {
		a := int(word[i] - 'a')
		b := int(word[i+1] - 'a')
		if abs(pos[a]-pos[b]) > 1 {
			return false
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	C := []int{7, 10, 4}
	words := []string{
		"abacaba",
		"cba",
		"db",
	}
	expect := "hjkildbacefg"
	runSample(t, C, words, expect)
}

func TestSample2(t *testing.T) {
	C := []int{100}
	words := []string{
		"abca",
	}
	expect := "abcdefghijkl"
	runSample(t, C, words, expect)
}
