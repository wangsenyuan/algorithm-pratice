package p2514

import "testing"

func runSample(t *testing.T, divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int, expect int) {
	res := minimizeSet(divisor1, divisor2, uniqueCnt1, uniqueCnt2)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	divisor1 := 12
	divisor2 := 3
	uniqueCnt1 := 2
	uniqueCnt2 := 10
	expect := 14
	runSample(t, divisor1, divisor2, uniqueCnt1, uniqueCnt2, expect)
}

func TestSample2(t *testing.T) {
	divisor1 := 5
	divisor2 := 2
	uniqueCnt1 := 2
	uniqueCnt2 := 20
	expect := 39
	runSample(t, divisor1, divisor2, uniqueCnt1, uniqueCnt2, expect)
}

func TestSample3(t *testing.T) {
	divisor1 := 2
	divisor2 := 5
	uniqueCnt1 := 683651932
	uniqueCnt2 := 161878530
	expect := 1367303863
	runSample(t, divisor1, divisor2, uniqueCnt1, uniqueCnt2, expect)
}
