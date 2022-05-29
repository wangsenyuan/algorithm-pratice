package p2287

import "testing"

func runSample(t *testing.T, s string, discount int, expect string) {
	res := discountPrices(s, discount)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1 2 $3 4 $5 $6 7 8$ $9 $10$"
	dis := 100
	expect := "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"
	runSample(t, s, dis, expect)
}

func TestSample2(t *testing.T) {
	s := "$76111 ab $6 $"
	dis := 48
	expect := "$39577.72 ab $3.12 $"
	runSample(t, s, dis, expect)
}
