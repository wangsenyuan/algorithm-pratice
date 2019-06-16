package p739

import (
	"reflect"
	"testing"
)

func TestSample(t *testing.T) {
	temps := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(temps)
	expect := []int{1, 1, 4, 2, 1, 1, 0, 0}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("daily temperatures %v, the warmer sequence should be %v, but got %v", temps, expect, res)
	}
}
