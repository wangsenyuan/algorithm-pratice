package p1291

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, low int, high int, expect []int) {
	res := sequentialDigits(low, high)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d %d, expect %v, but got %v", low, high, expect, res)
	}
}

func TestSample1(t *testing.T) {
	expect := []int{123, 234}
	runSample(t, 100, 300, expect)
}
