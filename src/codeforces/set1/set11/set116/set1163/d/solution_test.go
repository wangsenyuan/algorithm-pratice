package main

import "testing"

func runSample(t *testing.T, c string, s string, x string, expect int) {
	res := solve(c, s, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	c := "*****"
	s := "katie"
	x := "shiro"
	expect := 1
	runSample(t, c, s, x, expect)
}

func TestSample2(t *testing.T) {
	c := "caat"
	s := "caat"
	x := "a"
	expect := -1
	runSample(t, c, s, x, expect)
}

func TestSample3(t *testing.T) {
	c := "*a*"
	s := "bbb"
	x := "b"
	expect := 0
	runSample(t, c, s, x, expect)
}

func TestSample4(t *testing.T) {
	c := "***"
	s := "cc"
	x := "z"
	expect := 2
	runSample(t, c, s, x, expect)
}

func TestSample5(t *testing.T) {
	c := "ihllkwwcucecbjhs*wbgqgwjx*qfcgrohorzkxzjmjhksijtvmccvjtkjs*ubmbbeb*xxscizoglpdrq*ooqygqcieat*mcasbuym*i*vadlvbwplbosfmtzuutpthalrhzlqh*dldga*fdtpxpllqryewkbnhwzhfplwlnjpayxhhhxdwopctj*lg**didjgsvbsxrlugumtfzciozwgbaqegnkhckake*wckqrtlzctncuqljgeunmzrxanujqfjqdmyraplsosl*jakvnurbjgtoahwolvyni"
	s := "llv"
	x := "qscrflmsmhqyjppcu"
	expect := 0
	runSample(t, c, s, x, expect)
}
