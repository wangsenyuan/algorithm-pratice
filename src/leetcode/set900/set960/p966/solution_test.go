package p966

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, words []string, queries []string, expect []string) {
	res := spellchecker(words, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", words, queries, expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"KiTe", "kite", "hare", "Hare"}
	queries := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}
	expect := []string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"}
	runSample(t, words, queries, expect)
}
