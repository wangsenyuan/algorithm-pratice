package p763

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, S string, expect []int) {
	res := partitionLabels(S)
	if !reflect.DeepEqual(expect, res) {
		t.Errorf("sample %s, should give %v, but got %v", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "ababcbacadefegdehijhklij"
	expect := []int{9, 7, 8}
	runSample(t, S, expect)
}
