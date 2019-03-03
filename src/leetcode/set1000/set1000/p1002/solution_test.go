package p1002

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, A []string, expect []string) {
	res := commonChars(A)
	sort.Strings(res)
	sort.Strings(expect)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v, expect %v, but got %v", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []string{"bella", "label", "roller"}
	expect := []string{"e", "l", "l"}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []string{"cool", "lock", "cook"}
	expect := []string{"c", "o"}
	runSample(t, A, expect)
}
