package p986

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A, B []Interval, expect []Interval) {
	res := intervalIntersection(A, B)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %v, expect %v, but got %v", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []Interval{iv(0, 2), iv(5, 10), iv(13, 23), iv(24, 25)}
	B := []Interval{iv(1, 5), iv(8, 12), iv(15, 24), iv(25, 26)}
	expect := []Interval{iv(1, 2), iv(5, 5), iv(8, 10), iv(15, 23), iv(24, 24), iv(25, 25)}
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []Interval{iv(10, 12), iv(18, 19)}
	B := []Interval{iv(1, 6), iv(8, 11), iv(13, 17), iv(19, 20)}
	expect := []Interval{iv(10, 11), iv(19, 19)}
	runSample(t, A, B, expect)
}
