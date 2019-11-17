package p1258

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, synonyms [][]string, text string, expect []string) {
	res := generateSentences(synonyms, text)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %s, expect %v, but got %v", synonyms, text, expect, res)
	}
}

func TestSample1(t *testing.T) {
	synonyms := [][]string{
		{"happy", "joy"}, {"sad", "sorrow"}, {"joy", "cheerful"},
	}
	text := "I am happy today but was sad yesterday"

	expect := []string{
		"I am cheerful today but was sad yesterday",
		"I am cheerful today but was sorrow yesterday",
		"I am happy today but was sad yesterday",
		"I am happy today but was sorrow yesterday",
		"I am joy today but was sad yesterday",
		"I am joy today but was sorrow yesterday",
	}

	runSample(t, synonyms, text, expect)
}
