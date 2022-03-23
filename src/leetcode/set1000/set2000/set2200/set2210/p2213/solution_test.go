package p2213

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, queryCharacters string, queryIndices []int, expect []int) {
	res := longestRepeating(s, queryCharacters, queryIndices)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "babacc"
	query := "bcb"
	indices := []int{1, 3, 3}
	expect := []int{3, 3, 4}
	runSample(t, s, query, indices, expect)
}

func TestSample2(t *testing.T) {
	s := "abyzz"
	query := "aa"
	indices := []int{2, 1}
	expect := []int{2, 3}
	runSample(t, s, query, indices, expect)
}

func TestSample3(t *testing.T) {
	s := "geuqjmt"
	query := "bgemoegklm"
	indices := []int{3, 4, 2, 6, 5, 6, 5, 4, 3, 2}
	expect := []int{1, 1, 2, 2, 2, 2, 2, 2, 2, 1}
	runSample(t, s, query, indices, expect)
}

func TestSample4(t *testing.T) {
	s := "zlcbw"
	query := "ygwcnshib"
	indices := []int{3, 1, 3, 3, 4, 0, 3, 0, 0}
	expect := []int{1, 1, 2, 2, 2, 2, 1, 1, 1}
	runSample(t, s, query, indices, expect)
}

func TestSample5(t *testing.T) {
	s := "lptxlmmlcvfihnsulmomex"
	query := "tsqhthtruwlmbsoyyvjrhutjjwgbobkqbvihtspheychbtrbcxal"
	indices := []int{1, 5, 5, 2, 1, 15, 12, 0, 13, 5, 9, 3, 2, 1, 13, 7, 19, 21, 18, 19, 16, 9, 0, 2, 18, 8, 9, 10, 12, 10, 8, 5, 13, 2, 14, 0, 9, 3, 21, 0, 14, 5, 9, 16, 16, 5, 19, 8, 15, 10, 19, 15}
	expect := []int{2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1}
	runSample(t, s, query, indices, expect)
}
