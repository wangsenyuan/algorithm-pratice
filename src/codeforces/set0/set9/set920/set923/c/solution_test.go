package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, P []int, expect []int) {
	res := solve(A, P)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{8, 4, 13}
	P := []int{17, 2, 7}
	expect := []int{10, 3, 28}
	runSample(t, A, P, expect)
}

func TestSample2(t *testing.T) {
	A := []int{12, 7, 87, 22, 11}
	P := []int{18, 39, 9, 12, 16}
	expect := []int{0, 14, 69, 6, 44}
	runSample(t, A, P, expect)
}

func TestSample3(t *testing.T) {
	A := []int{331415699, 278745619, 998190004, 423175621, 42983144, 166555524, 843586353, 802130100, 337889448, 685310951}
	P := []int{226011312, 266003835, 342809544, 504667531, 529814910, 684873393, 817026985, 844010788, 993949858, 1031395667}
	expect := []int{128965467, 243912600, 4281110, 112029883, 223689619, 76924724, 429589, 119397893, 613490433, 362863284}
	runSample(t, A, P, expect)
}
