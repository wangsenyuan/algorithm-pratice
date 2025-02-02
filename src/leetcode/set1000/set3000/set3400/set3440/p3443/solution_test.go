package p3443

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := minCostGoodCaption(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "cdcd"
	expect := "cccc"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aca"
	expect := "aaa"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "antwfdps"
	expect := "nnnnnppp"
	runSample(t, s, expect)
}