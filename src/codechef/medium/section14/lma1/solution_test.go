package main

import "testing"

func runSample(t *testing.T, a, p, q Frac, D int, expect string) {
	res := solve(a, p, q, D)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := NewFrac(62, 100)
	p := NewFrac(271, 100)
	q := NewFrac(314, 100)
	D := 42
	expect := "17.903573407202216066481994459833795013850415"
	runSample(t, a, p, q, D, expect)
}

func TestSample2(t *testing.T) {
	a := NewFrac(91, 100)
	p := NewFrac(470, 100)
	q := NewFrac(667, 100)
	D := 18
	expect := "796.867901234567901234"
	runSample(t, a, p, q, D, expect)
}

func TestSample3(t *testing.T) {
	a := NewFrac(1, 100)
	p := NewFrac(981, 100)
	q := NewFrac(863, 100)
	D := 66
	expect := "0.187143148658300173451688603203754718906234057749209264360779512294"
	runSample(t, a, p, q, D, expect)
}
