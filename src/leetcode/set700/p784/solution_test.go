package p784

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, S string, expect []string) {
	res := letterCasePermutation(S)
	sort.Strings(expect)
	sort.Strings(res)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %s, expect %v, but got %v", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "a1b2"
	expect := []string{"a1b2", "a1B2", "A1b2", "A1B2"}
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "3z4"
	expect := []string{"3z4", "3Z4"}
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := "12345"
	expect := []string{"12345"}
	runSample(t, S, expect)
}

func TestSample4(t *testing.T) {
	S := "TB"
	expect := []string{"tb", "tB", "Tb", "TB"}
	runSample(t, S, expect)
}
