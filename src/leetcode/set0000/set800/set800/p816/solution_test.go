package p816

import (
	"reflect"
	"sort"
	"testing"
)

func runSample(t *testing.T, S string, expect []string) {
	res := ambiguousCoordinates(S)
	sort.Strings(res)
	sort.Strings(expect)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %s, expect %v, but got %v", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "(123)"
	expect := []string{"(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"}
	runSample(t, S, expect)
}

func TestSample2(t *testing.T) {
	S := "(00011)"
	expect := []string{"(0.001, 1)", "(0, 0.011)"}
	runSample(t, S, expect)
}

func TestSample3(t *testing.T) {
	S := "(0123)"
	expect := []string{"(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"}
	runSample(t, S, expect)
}

func TestSample4(t *testing.T) {
	S := "(100)"
	expect := []string{"(10, 0)"}
	runSample(t, S, expect)
}
