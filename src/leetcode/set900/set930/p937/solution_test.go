package p937

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, logs []string, expect []string) {
	res := reorderLogFiles(logs)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", logs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	logs := []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}
	expect := []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"}
	runSample(t, logs, expect)
}
