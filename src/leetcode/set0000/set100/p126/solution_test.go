package p126

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, begin, end string, words []string, expect [][]string) {
	res := findLadders(begin, end, words)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	begin := "hit"
	end := "cog"
	words := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	expect := [][]string{
		{"hit", "hot", "dot", "dog", "cog"},
		{"hit", "hot", "lot", "log", "cog"},
	}
	runSample(t, begin, end, words, expect)
}
