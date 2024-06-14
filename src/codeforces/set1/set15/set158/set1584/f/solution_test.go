package main

import "testing"

func runSample(t *testing.T, words []string, expect string) {
	res := solve(words)

	if len(res) != len(expect) {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	for _, word := range words {
		for i, j := 0, 0; i < len(res); i++ {
			for j < len(word) && res[i] != word[j] {
				j++
			}
			if j == len(word) {
				t.Fatalf("Sample result %s, not found in %s", res, word)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	words := []string{
		"ABC",
		"CBA",
	}
	expect := "A"
	runSample(t, words, expect)
}

func TestSample2(t *testing.T) {
	words := []string{
		"bacab",
		"defed",
	}
	expect := ""
	runSample(t, words, expect)
}

func TestSample3(t *testing.T) {
	words := []string{
		"abcde",
		"aBcDe",
		"ace",
	}
	expect := "ace"
	runSample(t, words, expect)
}

func TestSample4(t *testing.T) {
	words := []string{
		"codeforces",
		"technocup",
	}
	expect := "coc"
	runSample(t, words, expect)
}
