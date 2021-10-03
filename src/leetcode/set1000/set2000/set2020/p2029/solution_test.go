package p2029

import "testing"

func runSample(t *testing.T, s string, k int, letter byte, repetition int, expect string) {
	res := smallestSubsequence(s, k, letter, repetition)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "leet"
	k := 3
	var letter byte = 'e'
	repetition := 1
	expect := "eet"
	runSample(t, s, k, letter, repetition, expect)
}

func TestSample2(t *testing.T) {
	s := "leetcode"
	k := 4
	var letter byte = 'e'
	repetition := 2
	expect := "ecde"
	runSample(t, s, k, letter, repetition, expect)
}

func TestSample3(t *testing.T) {
	s := "bb"
	k := 2
	var letter byte = 'b'
	repetition := 2
	expect := "bb"
	runSample(t, s, k, letter, repetition, expect)
}

func TestSample4(t *testing.T) {
	s := "aaabbbcccddd"
	k := 3
	var letter byte = 'b'
	repetition := 2
	expect := "abb"
	runSample(t, s, k, letter, repetition, expect)
}

func TestSample5(t *testing.T) {
	s := "mmmxmxymmm"
	k := 8
	var letter byte = 'm'
	repetition := 4
	expect := "mmmmxmmm"
	runSample(t, s, k, letter, repetition, expect)
}

func TestSample6(t *testing.T) {
	s := "abbdglhoquvvwwyzz"
	k := 15
	var letter byte = 'l'
	repetition := 1
	expect := "abbdglhoquvvwwy"
	runSample(t, s, k, letter, repetition, expect)
}

func TestSample7(t *testing.T) {
	s := "cdjjmnqqrrvwwwxyydvrqqqnhged"
	k := 21
	var letter byte = 'd'
	repetition := 2
	expect := "cdjjmnqqrrdvrqqqnhged"
	runSample(t, s, k, letter, repetition, expect)
}
