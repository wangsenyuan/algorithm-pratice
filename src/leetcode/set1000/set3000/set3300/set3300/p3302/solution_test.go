package p3302

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, word1 string, word2 string, expect []int) {
	res := validSequence(word1, word2)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	word1 := "vbcca"
	word2 := "abc"
	expect := []int{0, 1, 2}
	runSample(t, word1, word2, expect)
}

func TestSample2(t *testing.T) {
	word1 := "bacdc"
	word2 := "abc"
	expect := []int{1, 2, 4}
	runSample(t, word1, word2, expect)
}

func TestSample3(t *testing.T) {
	word1 := "aaaaaa"
	word2 := "aaabc"
	// expect := []int{1, 2, 4}
	runSample(t, word1, word2, nil)
}

func TestSample4(t *testing.T) {
	word1 := "abc"
	word2 := "abc"
	expect := []int{0, 1, 2}
	runSample(t, word1, word2, expect)
}
