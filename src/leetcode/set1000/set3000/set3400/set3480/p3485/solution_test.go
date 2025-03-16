package p3485

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, words []string, k int, expect []int) {
	res := longestCommonPrefix(words, k)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"jump", "run", "run", "jump", "run"}
	k := 2
	expect := []int{3, 4, 4, 3, 4}
	runSample(t, words, k, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"dog", "racer", "car"}
	k := 2
	expect := []int{0, 0, 0}
	runSample(t, words, k, expect)
}

func TestSample3(t *testing.T) {
	words := []string{"cdb", "c", "aee", "afdd", "dad", "bdebb", "cdecf", "a", "efdb", "cffe", "bed", "ba"}
	k := 2
	expect := []int{1, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2}
	runSample(t, words, k, expect)
}

func TestSample4(t *testing.T) {
	words := []string{"bfee", "fe", "cfd", "bcdf", "afb", "abb", "fe"}
	k := 1
	expect := []int{4, 4, 4, 4, 4, 4, 4}
	runSample(t, words, k, expect)
}
