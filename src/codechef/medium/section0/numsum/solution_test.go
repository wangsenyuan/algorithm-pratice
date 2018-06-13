package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "3", "6")
}

func TestSample2(t *testing.T) {
	runSample(t, "5", "15")
}

func TestSample3(t *testing.T) {
	runSample(t, "10", "55")
}

func TestSample4(t *testing.T) {
	runSample(t, "100", "5050")
}

func TestSample5(t *testing.T) {
	runSample(t, "1000", calculate("1000"))
}

func TestSample6(t *testing.T) {
	runSample(t, "10000", calculate("10000"))
}

func TestSample7(t *testing.T) {
	runSample(t, "100000", calculate("100000"))
}

func TestSample8(t *testing.T) {
	runSample(t, "1000000", calculate("1000000"))
}

func runAdd(t *testing.T, X, Y string, expect string) {
	x := toNum(X)
	y := toNum(Y)
	z := add(OP{x, 1, len(x)}, OP{y, 1, len(y)})
	res := toStr(z.rep)
	if res != expect {
		t.Errorf("Sample %s %s, expect %s, but got %s", X, Y, expect, res)
	}
}

func TestAdd1(t *testing.T) {
	runAdd(t, "100", "3010", "3110")
}

func TestAdd2(t *testing.T) {
	runAdd(t, "100", "910", "1010")
}

func TestAdd3(t *testing.T) {
	runAdd(t, "10000000", "910", "10000910")
}

func runSub(t *testing.T, X, Y string, expect string) {
	var x OP

	if X[0] == '-' {
		x = OP{toNum(X[1:]), -1, len(X) - 1}
	} else {
		x = OP{toNum(X), 1, len(X)}
	}
	var y OP
	if Y[0] == '-' {
		y = OP{toNum(Y[1:]), -1, len(Y) - 1}
	} else {
		y = OP{toNum(Y), 1, len(Y)}
	}

	z := sub(x, y)
	res := toStr(z.rep)
	if z.sign < 0 {
		res = "-" + res
	}
	if res != expect {
		t.Errorf("Sample %s %s, expect %s, but got %s", X, Y, expect, res)
	}
}

func TestSub1(t *testing.T) {
	runSub(t, "10", "2", "8")
}

func TestSub2(t *testing.T) {
	runSub(t, "2", "10", "-8")
}

func TestSub3(t *testing.T) {
	runSub(t, "10", "10", "0")
}

func TestSub4(t *testing.T) {
	runSub(t, "-10", "10", "-20")
}

func TestSub5(t *testing.T) {
	runSub(t, "-2", "-10", "8")
}
