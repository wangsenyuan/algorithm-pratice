package p2904

import "testing"

func runSample(t *testing.T, s string, k int, expect string) {
	res := shortestBeautifulSubstring(s, k)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "100011001"
	k := 3
	expect := "11001"
	runSample(t, s, k, expect)
}

func TestSample2(t *testing.T) {
	s := "1011"
	k := 2
	expect := "11"
	runSample(t, s, k, expect)
}

func TestSample3(t *testing.T) {
	s := "000"
	k := 1
	expect := ""
	runSample(t, s, k, expect)
}
