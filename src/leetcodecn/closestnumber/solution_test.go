package closestnumber

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, num int, expect []int) {
	res := findClosedNumbers(num)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %d, expect %v, but got %v", num, expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := 10023
	expect := []int{10027, 10014}
	runSample(t, num, expect)
}

func TestSample2(t *testing.T) {
	num := 1156403390
	expect := []int{1156403407, 1156403389}
	runSample(t, num, expect)
}
