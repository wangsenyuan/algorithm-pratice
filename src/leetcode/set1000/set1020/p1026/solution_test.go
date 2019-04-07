package p1026

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, queries []string, pattern string, expect []bool) {
	res := camelMatch(queries, pattern)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %s, expect %v, but got %v", queries, pattern, expect, res)
	}
}

func TestSample1(t *testing.T) {
	queries := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern := "FB"
	expect := []bool{true, false, true, true, false}
	runSample(t, queries, pattern, expect)
}

func TestSample2(t *testing.T) {
	queries := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern := "FoBa"
	expect := []bool{true, false, true, false, false}
	runSample(t, queries, pattern, expect)
}

func TestSample3(t *testing.T) {
	queries := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern := "FoBaT"
	expect := []bool{false, true, false, false, false}
	runSample(t, queries, pattern, expect)
}
