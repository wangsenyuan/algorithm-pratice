package p821

import (
	"testing"
	"reflect"
)

func runSample(t *testing.T, S string, C byte, expect []int) {
	res := shortestToChar(S, C)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %s %v, expect %v but got %v", S, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	S := "loveleetcode"
	expect := []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}
	runSample(t, S, 'e', expect)
}
