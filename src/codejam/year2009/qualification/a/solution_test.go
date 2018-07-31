package main

import (
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, L int, D int, N int, words []string, puzzles []string, expect []int) {
	res := solve(L, D, N, words, puzzles)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d %d %v %v, expect %v, but got %v", L, D, N, words, puzzles, expect, res)
	}
}

func TestSample1(t *testing.T) {
	L, D, N := 3, 5, 4
	S := `
	abc
	bca
	dac
	dbc
	cba
	`
	words := strings.FieldsFunc(S, func(x rune) bool { return x == ' ' || x == '\n' || x == '\t' })
	W := `
	(ab)(bc)(ca)
abc
(abc)(abc)(abc)
(zyx)bc
	`
	puzzles := strings.FieldsFunc(W, func(x rune) bool { return x == ' ' || x == '\n' || x == '\t' })
	expect := []int{2, 1, 3, 0}
	runSample(t, L, D, N, words, puzzles, expect)
}
